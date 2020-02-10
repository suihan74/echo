package main

import (
    "fmt"
    "log"
    "strings"
    "net/http"

    // REST
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"

    // websocket
    "github.com/gorilla/websocket"

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
        // クライアントから送られてきた JWT 取得
        authHeader := r.Header.Get("Authorization")
        jwt := strings.Replace(authHeader, "Bearer ", "", 1)

        // この部分キャッシュしたい
        uid, err := verifyToken(jwt)
        if err != nil {
            fmt.Printf("error verifying ID token: %v\n", err)
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("error verifying ID token\n"))
            return
        }

        // ユーザー情報を登録する
        user := registerUser(uid)

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
var wsBroadcast = make(chan WebSocketMessage)
var wsUpgrader = websocket.Upgrader{ CheckOrigin: func(r *http.Request) bool { return true } }

/** WebSocketひらく */
func handleWebSocketClients(w http.ResponseWriter, r *http.Request) {
    // TODO: 認証情報を扱うようにする

    go broadcastMessagesToWebSocketClients()
    websocket, err := wsUpgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal("error upgrading GET request to a websocket::", err)
    }
    defer websocket.Close()

    wsClients[websocket] = true

    for {
        var message WebSocketMessage
        err := websocket.ReadJSON(&message)
        if err != nil {
            log.Printf("error occurred while reading message: %v\n", err)
            delete(wsClients, websocket)
            break
        }
        wsBroadcast <- message
    }
}

/** WebSocketのメッセージがbroadcastされてきたときに接続中のクライアントにメッセージを配信する */
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

    // 認証
    router.HandleFunc("/auth", authMiddleware(signIn)).Methods("GET")

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
    router.HandleFunc("/socket", handleWebSocketClients)

////// エンドポイントここまで //////

    log.Fatal(http.ListenAndServe(port, handlers.CORS(/*allowedOrigins,*/ allowedMethods, allowedHeaders)(router)))
}
