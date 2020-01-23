# echo

完全匿名SNSをつくってみる（つくってみたい）

## やりたいこと (箇条書き)

- 完全匿名なSNS

- ユーザー名は表示されない

- リプライやRTを気にしなくていい（ない）

- 自分が書いた投稿はまとめて見られる

- 引用はできるようにする。ただし引用されても通知はされない

- fav。自分がfavつけた投稿はまとめて見られる。

- ユーザーミュート。不快な投稿をしたユーザーをミュートする。

## 作業メモ

web系ほぼ知らないので勉強しながらという感じなのでアレです。

- Vue.jsでフロントつくる

- Goでバックエンドつくる

- Firebaseかどっかにデプロイする

## ローカルで動かすやつ

1. `sudo service postgresql start`する

2. /backで`go run *.go`する

3. /fontで`npm run dev`する

4. localhost:8080を開く

## 参考

### Go

- [go modの使い方](https://blog.mmmcorp.co.jp/blog/2019/10/10/go-mod/)

### js

- [nodebrewをインストール](https://contents.shinonomekazan.com/tips/wsl-with-node.html#node-jsのインストール)

- [Vue.js + Go + Firebaseでwebアプリを作る](https://qiita.com/po3rin/items/d3e016d01162e9d9de80)

- [npmコマンドの使い方](https://qiita.com/wifecooky/items/c3be77e54233fcfca376)

