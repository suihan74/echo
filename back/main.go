package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "strings"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"

    firebase "firebase.google.com/go"
    "google.golang.org/api/option"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

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
}

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

func main() {
    initDatabase()

    allowedOrigins := handlers.AllowedOrigins([]string {"http://localhost:8080"})
    allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
    allowedHeaders := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})

    router := mux.NewRouter()

////// エンドポイント //////

    // 投稿関連
    router.HandleFunc("/post", authMiddleware(postEndPoint)).Methods("POST")
    router.HandleFunc("/posts", authMiddleware(getPostsEndPoint)).Methods("GET")

    // ユーザー関連
    router.HandleFunc("/user", authMiddleware(getUserEndPoint)).Methods("GET")

////// エンドポイントここまで //////

    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}
