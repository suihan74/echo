# echo

完全匿名SNSをつくってみる（つくってみたい）

## いまこんなの

![現況](https://github.com/suihan74/echo/wiki/images/ss_3.png "2020/02/02")

### 2020/02/02

- Fav機能を作り始めた

    - 同一の投稿に複数回favできてしまう(一回だけにするかこのままにするか)

- 投稿APIの投げ方を変えた

    - 前: Post型から変換したJSONを投げる
    
    - 今: クエリパラメータに必要なものを書く e.g. /post?text=hoge&quote_id=1234

- 投稿ボタンのスタイル適当に書いた

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

1. `sudo service postgresql start`する

2. /backで`go run *.go`する

3. /frontで`npm run dev`する

4. localhost:8080を開く

## 参考

### Go

- [go modの使い方](https://blog.mmmcorp.co.jp/blog/2019/10/10/go-mod/)

- [gormクエリの使い方](http://gorm.io/ja_JP/docs/query.html)

- [http status code](http://golang.jp/pkg/http)

### js

- [nodebrewをインストール](https://contents.shinonomekazan.com/tips/wsl-with-node.html#node-jsのインストール)

- [Vue.js + Go + Firebaseでwebアプリを作る](https://qiita.com/po3rin/items/d3e016d01162e9d9de80)

- [npmコマンドの使い方](https://qiita.com/wifecooky/items/c3be77e54233fcfca376)

### styles

- [icons](https://freedesignresources.net/100-free-minimal-line-icons/)
