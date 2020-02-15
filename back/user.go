package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	// firebase
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

////////////////////////////////////////////////////////////

// User ユーザー情報
type User struct {
	// ユーザーID
	ID int64 `json:"id" gorm:"primary_key"`

	// ユーザートークン
	Token string `json:"-" gorm:"primary_key"`

	// 投稿数
	PostsCount int64 `json:"posts_count"`

	// Fav数
	FavoritesCount int64 `json:"favorites_count"`

	// 被Fav数
	FavoritedCount int64 `json:"favorited_count"`
}

// AuthResult 認証レスポンス
type AuthResult struct {
	Succeeded bool `json:"succeeded"`
}

////////////////////////////////////////////////////////////

// firebaseのJWTを検証し、UIDを取得する
func verifyToken(jwt string) (string, error) {
	// Firebase SDK のセットアップ
	opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", fmt.Errorf("environment error")
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", fmt.Errorf("environment error")
	}

	// JWT 検証
	token, err := auth.VerifyIDTokenAndCheckRevoked(context.Background(), jwt)
	if err != nil {
		return "", err
	}

	return token.UID, nil
}

func signIn(w http.ResponseWriter, r *http.Request, db *gorm.DB, user User) {
	response := AuthResult{Succeeded: true}
	json.NewEncoder(w).Encode(response)
}

////////////////////////////////////////////////////////////

/*
 * ユーザー情報を取得する
 */
func getUserEndPoint(w http.ResponseWriter, r *http.Request, db *gorm.DB, user User) {
}
