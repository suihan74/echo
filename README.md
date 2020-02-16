# echo

完全匿名SNSをつくってみる（つくってみたい）

## いまこんなの

![現況](https://github.com/suihan74/echo/wiki/images/ss_3.png "2020/02/02")

### 2020/02/15

- WebSocketの接続open時にユーザー認証するようにした

    - とりあえずopen時にjwt投げるようにした
    
    - リアルタイム配信されるPost情報を配信先のユーザーごとに修正するようにできた

### 2020/02/10

- WebSocketでTLをリアルタイム更新するようにした

    - 自分の投稿を判別するための処理をまだ作っていないのでその点RESTと比して問題あり

    - wsにはヘッダを付加できない？（代替手法はある様子）のでそこをバック側で対応する必要
    
- 認証が必要なAPI使用時に毎回JWT検証しているので時間がかかっている

    - JWTとUIDのペアをキャッシュした方がよさそう。go-cacheなど

## やりたいこと (箇条書き)

- 完全匿名なSNS

    - ユーザー名は表示されない

    - 発言の価値が「誰の発言か」に依らない空間を作りたい

- リプライやRTを気にしなくていい（ない）

    - 匿名前提の空間で他人と積極的に関わっても碌なことはない

    - 引用はできるようにする。ただし引用されても通知はされない

- 自分が書いた投稿はまとめて見られる

- fav。自分がfavつけた投稿はまとめて見られる

    - 投稿に付けられたfav数は自分の投稿に限り確認することができる

- ワードミュート + ユーザーミュート。不快な投稿をしたユーザーをミュートする

## 作業メモ

web系ほぼ知らないので勉強しながらという感じなのでアレです。

- Vue.jsでフロントつくる

- Goでバックエンドつくる

- PostgreSQLつかう

  - Goでやるにはgormとか使う

- Firebaseかどっかにデプロイする

  - とりあえず認証機能はFirebaseで提供されているやつを使うことにした

  - バックエンド込みでの置き方よくわからない。どうしよう

## ローカルで動かすやつ

### 実行環境構築手順

[echoの実行環境構築メモ - すいはんぶろぐ.io](https://suihan74.github.io/posts/2020/02_09_00_web_dev_installation_tips/)

### 実行手順

1. `sudo service postgresql start`する

2. /backで`go run *.go`する

3. /frontで`npm run dev`する

4. localhost:8080を開く

## 参考

### Go

- [go modの使い方](https://blog.mmmcorp.co.jp/blog/2019/10/10/go-mod/)

- [gormクエリの使い方](http://gorm.io/ja_JP/docs/query.html)

- [http status code](http://golang.jp/pkg/http)

- [websocket](https://qiita.com/__init__/items/08cbc3a870178fd6fc32)

### js

- [nodebrewをインストール](https://contents.shinonomekazan.com/tips/wsl-with-node.html#node-jsのインストール)

- [Vue.js + Go + Firebaseでwebアプリを作る](https://qiita.com/po3rin/items/d3e016d01162e9d9de80)

- [npmコマンドの使い方](https://qiita.com/wifecooky/items/c3be77e54233fcfca376)

### styles

- [icons](https://freedesignresources.net/100-free-minimal-line-icons/)
