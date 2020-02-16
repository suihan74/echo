package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	// REST
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

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

// AuthHandlerFunc 要認証APIのハンドラ
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
func registerUser(userID string) User {
	var user User
	db.Where(User{Token: userID}).Find(&user)
	if user.ID == 0 {
		user.Token = userID
		db.Create(&user)
	}
	return user
}

////////////////////////////////////////////////////////////
// websocket

// WebSocketMessageType wsで渡されるPostに何が起こったかを表す
type WebSocketMessageType int

// enum: WebSocketMessageType
const (
	CREATE WebSocketMessageType = iota
	UPDATE
	DELETE
)

// WebSocketMessage websocket用メッセージ
type WebSocketMessage struct {
	Type WebSocketMessageType `json:"type"`
	Post Post                 `json:"post"`
}

// WebSocketAuthMessage websocketユーザー認証用メッセージ
type WebSocketAuthMessage struct {
	Token string `json:"token"`
}

var wsClients = make(map[*websocket.Conn]User)
var wsBroadcast = make(chan WebSocketMessage)
var wsUpgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

/** WebSocketひらく */
func handleWebSocketClients(w http.ResponseWriter, r *http.Request) {
	go broadcastMessagesToWebSocketClients()
	client, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("error upgrading GET request to a websocket::", err)
	}
	defer client.Close()

	// ユーザー認証を待機する
	for {
		var message WebSocketAuthMessage
		err := client.ReadJSON(&message)
		if err != nil {
			log.Printf("error occurred while authorization: %v\n", err)
			delete(wsClients, client)
			return
		}

		uid, err := verifyToken(message.Token)
		if err != nil {
			fmt.Printf("error verifying ID token: %v\n", err)
			delete(wsClients, client)
			return
		}

		// ユーザー情報を登録する
		user := registerUser(uid)
		wsClients[client] = user
		break
	}

	// 通信終了を待機
	for {
		_, _, err := client.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v\n", err)
			}
			break
		}
	}
	delete(wsClients, client)
}

/** WebSocketのメッセージがbroadcastされてきたときに接続中のクライアントにメッセージを配信する */
func broadcastMessagesToWebSocketClients() {
	for {
		message := <-wsBroadcast

		// 投稿ユーザーに配信するPost
		postForUser := message.Post
		postForUser.IsYours = true
		msgForUser := WebSocketMessage{Type: message.Type, Post: postForUser}

		// 投稿ユーザー以外のユーザーに配信するPost
		postForOthers := message.Post
		postForOthers.IsYours = false
		postForOthers.FavoritedCount = 0
		msgForOthers := WebSocketMessage{Type: message.Type, Post: postForOthers}

		msg := map[bool]WebSocketMessage{true: msgForUser, false: msgForOthers}

		for client, user := range wsClients {
			err := client.WriteJSON(msg[user.ID == message.Post.UserID])
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
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})

	router := mux.NewRouter()

	////// エンドポイント //////

	// 認証
	router.HandleFunc("/auth", authMiddleware(signIn)).Methods("GET")

	// 投稿関連
	router.HandleFunc("/post", authMiddleware(postEndPoint)).Methods("POST")
	router.HandleFunc("/post", authMiddleware(deletePostEndPoint)).Methods("DELETE")
	router.HandleFunc("/post", authMiddleware(getPostEndPoint)).Methods("GET")
	router.HandleFunc("/posts", authMiddleware(getPostsEndPoint)).Methods("GET")

	router.HandleFunc("/quotes", authMiddleware(getQuotesEndPoint)).Methods("GET")

	// Fav関連
	router.HandleFunc("/favs", authMiddleware(getFavoritesEndPoint)).Methods("GET")
	router.HandleFunc("/fav", authMiddleware(favoritePostEndPoint)).Methods("POST")

	// ユーザー関連
	router.HandleFunc("/user", authMiddleware(getUserEndPoint)).Methods("GET")

	// WebSocket接続
	router.HandleFunc("/socket", handleWebSocketClients)

	////// エンドポイントここまで //////

	log.Fatal(http.ListenAndServe(port, handlers.CORS( /*allowedOrigins,*/ allowedMethods, allowedHeaders)(router)))
}
