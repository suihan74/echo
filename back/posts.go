package main

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

/*
 * 投稿内容
 */
type Post struct {
    /** 投稿ID */
    Id int64 `json:"Id" gorm:"primary_key"`
    /** 投稿ユーザー */
    UserId int64 `json:"-" gorm:"primary_key" sql:"not null"`
    /** 投稿内容 */
    Text string `json:"Text" sql:"not null"`
    /** 投稿時刻 */
    Timestamp int64 `json:"Timestamp" sql:"not null"`

    /** 引用対象ID nullable */
    QuoteId int64 `json:"QuoteId"`
    /** 引用対象ポスト */
    QuotePost *Post `json:"QuotePost" gorm:"-"`
}

/*
 * 投稿を受け付ける
 */
func postEndPoint(w http.ResponseWriter, r *http.Request, db *gorm.DB, user User) {
    var post Post
    json.NewDecoder(r.Body).Decode(&post)

    // DBに登録
    post.UserId = user.Id
    post.Timestamp = time.Now().Unix()
    db.Create(&post)

    // 引用先を探す
    if post.QuoteId != 0 {
        var quote Post
        quote.Id = post.QuoteId
        db.Where(quote).Find(&quote)
        post.QuotePost = &quote
    }

    // 結果を返す
    json.NewEncoder(w).Encode(post)
}

/*
 * 最近の投稿を取得する
 */
func getPostsEndPoint(w http.ResponseWriter, r *http.Request, db *gorm.DB, user User) {
    // 最新の投稿をN個まで取得する
    var posts []Post
    db.Order("id desc").Limit(20).Find(&posts)

    json.NewEncoder(w).Encode(posts)
}
