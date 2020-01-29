package main

import (
    "encoding/json"
    "fmt"
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
    Id int64 `json:"id" gorm:"primary_key"`
    /** 投稿ユーザー */
    UserId int64 `json:"-" gorm:"primary_key" sql:"not null"`
    /** 投稿内容 */
    Text string `json:"text" sql:"not null"`
    /** 投稿時刻 */
    Timestamp int64 `json:"timestamp" sql:"not null"`

    /** ユーザーの投稿かを識別 */
    IsYours bool `json:"is_yours" gorm:"-"`

    /** 引用対象ID nullable */
    QuoteId int64 `json:"quote_id"`
    /** 引用対象ポスト */
    QuotePost *Post `json:"quote_post" gorm:"-"`
}

/**
 * 指定keyのクエリパラメータを取得する
 */
func getQueryParam(r *http.Request, key string) (string, error) {
    params, ok := r.URL.Query()[key]

    if !ok || len(params[0]) < 1 {
        return "", fmt.Errorf("no valid parameters: \"%s\"", key)
    }

    return params[0], nil
}

/**
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
        db.Where("id = ?", post.QuoteId).Find(&quote)
        post.QuotePost = &quote
    }

    post.IsYours = true

    // 結果を返す
    json.NewEncoder(w).Encode(post)
}

/**
 * 投稿を削除する
 *
 * DELETE /post?id=1234
 * response: 200(OK) or 400(BadRequest)
 */
func deletePostEndPoint(w http.ResponseWriter, r *http.Request, db *gorm.DB, user User) {
    postIdStr, err := getQueryParam(r, "id")
    if err != nil {
        fmt.Printf("error: cannot to delete a post")
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("deleting post failure\n"))
    }

    var post Post
    db.Find(&post, "user_id = ? AND id = ?", user.Id, postIdStr)
    if post.Id == 0 {
        fmt.Printf("error: cannot to delete a post")
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("deleting post failure\n"))
    }

    // DBから削除
    db.Delete(Post{}, "id = ?", postIdStr)

    // 結果を返す
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("deleting post succeeded\n"))
}

/**
 * 最近の投稿を取得する
 *
 * response: Post
 */
func getPostsEndPoint(w http.ResponseWriter, r *http.Request, db *gorm.DB, user User) {
    // 最新の投稿をN個まで取得する
    var posts []Post
    db.Order("id desc").Limit(20).Find(&posts)

    for idx, p := range posts {
        posts[idx].IsYours = p.UserId == user.Id
        if p.QuoteId != 0 {
            var quote Post
            db.Where("id = ?", p.QuoteId).First(&quote)
            if quote.Id != 0 {
                posts[idx].QuotePost = &quote
            }
        }
    }

    json.NewEncoder(w).Encode(posts)
}
