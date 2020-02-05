package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "strings"
    "net/http"

    // REST
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"

    // websocket
    "github.com/gorilla/websocket"

    // firebase
    firebase "firebase.google.com/go"
    "google.golang.org/api/option"

    // DB
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

////////////////////////////////////////////////////////////

var db *gorm.DB

func initDatabase() {
    var err error
    db, err = gorm.Open(
        "postgres",
        "user=echo dbname=echo password=echo sslmode=disable")
    if err != nil {
        panic(err)
    }

    db.AutoMigrate(&User{})
    db.AutoMigrate(&Post{})
    db.AutoMigrate(&Fav{})
}

////////////////////////////////////////////////////////////

type AuthHandlerFunc func(w http.ResponseWriter, r *http.Request, db *gorm.DB, user User)

/**
 * ミドルウェア
 *
 * Firebaseで認証する
 */
func authMiddleware(next AuthHandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Firebase SDK のセットアップ
        opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
        app, err := firebase.NewApp(context.Background(), nil, opt)
        if err != nil {
            fmt.Printf("error: %v\n", err)
            os.Exit(1)
        }

        auth, err := app.Auth(context.Background())
        if err != nil {
            fmt.Printf("error: %v\n", err)
            os.Exit(1)
        }

        // クライアントから送られてきた JWT 取得
        authHeader := r.Header.Get("Authorization")
        idToken := strings.Replace(authHeader, "Bearer ", "", 1)

        // JWT 検証
        token, err := auth.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
        if err != nil {
            if err.Error() == "ID token has been revoked" {
                fmt.Printf("ID token has been revoked: %v\n", err)
                w.WriteHeader(http.StatusNonAuthoritativeInfo)
                w.Write([]byte("error verifying ID token\n"))
            } else {
                fmt.Printf("error verifying ID token: %v\n", err)
                w.WriteHeader(http.StatusUnauthorized)
                w.Write([]byte("error verifying ID token\n"))
            }
            return
        }

        // ユーザー情報を登録する
        user := registerUser(token.UID)

        next(w, r, db, user)
    }
}

/*
 * ユーザー情報を(必要なら)登録する
 */
func registerUser(userId string) User {
    var user User
    db.Where(User{Token: userId}).Find(&user)
    if user.Id == 0 {
        user.Token = userId
        db.Create(&user)
    }
    return user
}

////////////////////////////////////////////////////////////
// websocket

var wsClients = make(map[*websocket.Conn]bool)
var wsBroadcast = make(chan Message)
var wsUpgrader = websocket.Upgrader{}

type Message struct {
    Message string `json:message`
}

func HandleWebSocketClients(w http.ResponseWriter, r *http.Request, db *gorm.DB, user User) {
    go broadcastMessagesToWebSocketClients()
    websocket, err := wsUpgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal("error uprading GET request to a websocket::", err)
    }
    defer websocket.Close()

    wsClients[websocket] = true

    for {
        var message Message
        err := websocket.ReadJSON(&message)
        if err != nil {
            log.Printf("error occurred while reading message: %v\n", err)
            delete(wsClients, websocket)
            break
        }
        wsBroadcast <- message
    }
}

func broadcastMessagesToWebSocketClients() {
    for {
        message := <- wsBroadcast
        for client := range wsClients {
            err := client.WriteJSON(message)
            if err != nil {
                log.Printf("error occured while writing message to client: %v\n", err)
                client.Close()
                delete(wsClients, client)
            }
        }
    }
}

////////////////////////////////////////////////////////////

func main() {
    port := ":8000"

    initDatabase()

//    allowedOrigins := handlers.AllowedOrigins([]string { "http://localhost:8080/" })
    allowedMethods := handlers.AllowedMethods([]string{ "GET", "POST", "DELETE", "PUT" })
    allowedHeaders := handlers.AllowedHeaders([]string{ "Authorization", "Content-Type" })

    router := mux.NewRouter()

////// エンドポイント //////

    // 投稿関連
    router.HandleFunc("/post", authMiddleware(postEndPoint)).Methods("POST")
    router.HandleFunc("/post", authMiddleware(deletePostEndPoint)).Methods("DELETE")
    router.HandleFunc("/post", authMiddleware(getPostEndPoint)).Methods("GET")
    router.HandleFunc("/posts", authMiddleware(getPostsEndPoint)).Methods("GET")

    // Fav関連
    router.HandleFunc("/fav", authMiddleware(favoritePostEndPoint)).Methods("POST")

    // ユーザー関連
    router.HandleFunc("/user", authMiddleware(getUserEndPoint)).Methods("GET")

    // WebSocket接続
    router.HandleFunc("/socket", authMiddleware(HandleWebSocketClients))

////// エンドポイントここまで //////

    log.Fatal(http.ListenAndServe(port, handlers.CORS(/*allowedOrigins,*/ allowedMethods, allowedHeaders)(router)))
}
