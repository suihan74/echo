package main

import (
    "net/http"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

////////////////////////////////////////////////////////////

/**
 * ユーザー情報
 */
type User struct {
    /** ユーザーID */
    Id int64 `json:"Id" gorm:"primary_key"`

    /** ユーザートークン */
    Token string `json:"-" gorm:"primary_key"`

    /** 投稿数 */
    PostsCount int64 `json:"posts_count"`

    /** Fav数 */
    FavoritesCount int64 `json:"favorites_count"`

    /** 被Fav数 */
    FavoritedCount int64 `json:"favorited_count"`
}

////////////////////////////////////////////////////////////

/*
 * ユーザー情報を取得する
 */
func getUserEndPoint(w http.ResponseWriter, r *http.Request, db *gorm.DB, user User) {
}
