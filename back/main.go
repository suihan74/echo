package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "strings"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"

    firebase "firebase.google.com/go"
    "google.golang.org/api/option"
)

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
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
        token, err := auth.VerifyIDToken(context.Background(), idToken)
        if err != nil {
            fmt.Printf("error verifying ID token: %v\n", err)
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("error verifying ID token\n"))
            return
        }

        log.Printf("Verified ID token: %v\n", token)
        next.ServeHTTP(w, r)
    }
}

func public(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello public!\n"))
}

func private(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello private!\n"))
}

type Post struct {
    Text string `json:\"text\"`
}

/*
 * 投稿を受け付ける
 */
func postEcho(w http.ResponseWriter, r *http.Request) {
    var post Post
    json.NewDecoder(r.Body).Decode(&post)

    // 結果を返す
    json.NewEncoder(w).Encode(post)

    fmt.Printf("post: %v\n", post.Text)
}

func main() {
    allowedOrigins := handlers.AllowedOrigins([]string {"http://localhost:8080"})
    allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
    allowedHeaders := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})

    router := mux.NewRouter()

    // エンドポイント
    router.HandleFunc("/public", public).Methods("GET")
    router.HandleFunc("/private", authMiddleware(private)).Methods("GET")
    router.HandleFunc("/posts/post", authMiddleware(postEcho)).Methods("POST")

    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}
