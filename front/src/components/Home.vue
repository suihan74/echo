<template>
  <div id='home'>
    <!-- 投稿エリア -->
    <div id="post-area">
      <div id="post-quote-area" v-if="quotePost">
        <div class="post-quote-body">
          <span class="quote-text">{{ quotePost.text }}</span>
          <time class="quote-timestamp">{{ timeToString(quotePost.timestamp) }}</time>
        </div>

        <button class="post-quote-close" @click="removeQuote">
          ×
        </button>
      </div>

      <textarea
        class="post-text"
        placeholder="いまどうしてる？"
        @keydown.prevent.ctrl.enter="post"
        v-model.trim="text"/>
      <!-- 投稿エリア下部 -->
      <div class="post-commands">
        <span class="counter">{{ text.length }} 文字</span>
        <button class="post-button" @click="post" title="Ctrl+Enter でも投稿できます">
          post
        </button>
      </div>
    </div>

    <div class="timeline">
      <ul class="timeline-items">
        <li class="timeline-item" v-for="post in posts" :key="post.id" @click="clickPost(post)">
          <span class="timeline-item-body">{{ post.text }}</span>
          <time class="timeline-item-timestamp">{{ timeToString(post.timestamp) }}</time>

          <!-- 引用情報 -->
          <div class="timeline-item-quote-area" v-if="post.quote_post">
            <span class="timeline-item-quote-text">{{ post.quote_post.text }}</span>
            <time class="timeline-item-quote-timestamp">{{ timeToString(post.quote_post.timestamp) }}</time>
          </div>

          <!-- コマンド -->
          <div class="timeline-item-commands">
            <!-- 引用 -->
            <div class="timeline-item-command">
              <img class="timeline-command-item" src="static/images/ic-message-type.svg" @click="quote(post)"/>
            </div>
            <!-- スター -->
            <div class="timeline-item-command">
              <img class="timeline-command-item" src="static/images/ic-star.svg" @click="favPost(post.id)"/>
              <span v-if="post.is_yours && post.favorited_count>0" class="timeline-command-item-favs-count">
                {{ post.favorited_count }}
              </span>
            </div>
            <!-- 消去 -->
            <div class="timeline-item-command">
              <img class="timeline-command-item" v-if="post.is_yours" src="static/images/ic-cross.svg" @click="deletePost(post.id)"/>
            </div>
          </div>
        </li>
      </ul>
    </div>

    <!--
    <div class="footer">
      <button @click="signOut">Sign out</button>
    </div>-->
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import axios from 'axios'
import { Service } from 'axios-middleware'
import Moment from 'moment'
import firebase from 'firebase/app'
import router from '../router'

// axiosに処理を挟み込む
const service = new Service(axios)
service.register({
  async onRequest (config) {
    config.baseURL = 'http://localhost:8000'
    const token = localStorage.getItem('jwt')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  }
})

@Component
export default class Home extends Vue {
  text: string = ''
  posts: Array<any> = []
  quotePost: any = null
  socket: WebSocket = null
  initialized: boolean = false

  /** 投稿する */
  post () {
    if (this.text.length === 0) {
      console.error('空文字列は投稿されない')
      return
    }

    const data = {
      text: this.text,
      quote_id: (this.quotePost ? this.quotePost.id : 0)
    }

    axios.post('/post', data).then(res => {
      // 投稿後編集内容をクリアしてタイムラインを更新する
      this.text = ''
      this.quotePost = null
    }).catch(err => {
      console.error(err)
    })
  }

  /** (自分の)投稿を削除する */
  deletePost (postId) {
    axios.delete('/post?id=' + postId).then(res => {
      console.log('deleted post: ' + postId)
    }).catch(err => {
      console.error(err)
    })
  }

  /** 最新の投稿を取得する */
  getPosts () {
    axios.get('posts').then(res => {
      this.posts = res.data
    }).catch(err => {
      console.error(err)
    })
  }

  /** 投稿をお気に入りにする */
  favPost (postId) {
    axios.post('/fav?id=' + postId).then(res => {
      console.log('favorite post: ' + postId)
    }).catch(err => {
      console.error(err)
    })
  }

  /** 引用対象を設定する */
  quote (q) {
    this.quotePost = q
    console.log('quote: ' + q.id)
  }

  /** 引用対象を除去する */
  removeQuote () {
    this.quotePost = null
  }

  clickPost (post) {
    console.log('id: ' + post.id)
    console.log('text: ' + post.text)
    console.log('quote_id: ' + post.quote_id)
    console.log('favs: ' + post.favorited_count)
    console.log('is_yours: ' + post.is_yours)
  }

  signOut () {
    firebase.auth().signOut().then(() => {
      localStorage.removeItem('jwt')
      router.push('/signin')
    })
  }

  /** タイムスタンプを表示用に加工する */
  timeToString (unixtime) {
    const date = new Date(unixtime * 1000)
    return Moment(date).format('YYYY-MM-DD HH:mm:ss')
  }

  /** 良い感じにPostをTLに挿入 */
  insertTimelineItem (post) {
    const idx = this.posts.findIndex(p => p.id === post.id)
    if (idx === -1) {
      // 追加
      const insertIdx = this.posts.findIndex(p => p.id < post.id)
      this.posts.splice(insertIdx, 0, post)
    } else {
      // 更新
      this.posts.splice(idx, 1, post)
    }
  }

  /** Postをタイムラインから削除 */
  removeTimelineItem (post) {
    const idx = this.posts.findIndex(p => p.id === post.id)
    if (idx !== -1) {
      this.posts.splice(idx, 1)
    }
  }

  // ロード後にタイムラインを取得する
  mounted () {
    if (!this.initialized) {
      this.initialized = true

      const inst = this
      // websocket
      const socket = new WebSocket('ws://localhost:8000/socket')
      socket.onopen = function (msg) {
        console.log('socket opened')
        socket.send(JSON.stringify({
          token: localStorage.getItem('jwt')
        }))
      }
      socket.onmessage = function (msg) {
        const data = JSON.parse(msg.data)
        if (data.type === 0 || data.type === 1) {
          // 追加・更新
          inst.insertTimelineItem(data.post)
        } else if (data.type === 2) {
          // 消去
          inst.removeTimelineItem(data.post)
        }
      }
      this.socket = socket

      // TL更新
      this.getPosts()
    }
  }
}
</script>

<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
button {
  width: auto;
  height: 42px;
  margin: 10px 0;
  padding: 10px;
}
textarea {
  resize: none;
}

/* 画面 */
#home {
  background: #fefefe;
  overflow-x: hidden;
  display: flex;
  flex-direction: column;
  margin: 0 3vw;
  max-width: 600px;
  margin: 0 auto;
}

/* 投稿エリア */
#post-area {
  background: #dfefef;
  display: flex;
  flex-direction: column;
  padding: 12px 16px;
}

/* 投稿オプションエリア */
.post-commands {
  padding: 8px 0;
  display: flex;
  flex-direction: row;
}

/* 編集中の文字列 */
.post-text {
  font-size: 14pt;
  height: 48px;
  padding: 12px;
  border-color: #aaa;
}
.post-text:hover {
  border-color: #668ad8;
}
.post-button {
  margin: 0;
  padding: 0 16px;
  background-color: #346add;
  color: #fff;
  border: none;
  border-radius: 12px;
  transition: .4s;
}
.post-button:hover {
  background-color: #668ad8;
}
/* 文字数カウンタ */
.counter {
  font-size: 12pt;
  margin-right: auto;
}

/* 投稿時引用情報 */
#post-quote-area {
  background: #acbcbc;
  padding: 6px 12px;
  display: flex;
  flex-direction: row;
}
.post-quote-body {
  display: flex;
  flex-direction: column;
  text-align: left;
  margin-right: auto;
}
.post-quote-close {
  display: inline-block;
  text-decoration: none;
  color: #668ad8;
  width: 28px;
  height: 28px;
  line-height: 0px;
  border-radius: 50%;
  border: solid 0px;
  text-align: center;
  overflow: hidden;
  font-weight: bold;
  transition: .4s;
}
.post-quote-close:hover {
  background: #b3e1ff;
}
.quote-text {
  color: #884444;
  font-size: 11pt;
}
.quote-timestamp {
  color: #666;
  font-size: 10pt;
}

.timeline {
  overflow: hidden;
  margin: 0;
}
.timeline-items {
  margin: 8px 0 0 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  _zoom: 1;
  overflow: hidden;
}
.timeline-item {
  padding: 10px 16px;
  text-align: left;
  display: flex;
  flex-direction: column;
  margin: -1px 0 0 0;
  border-top: 1px dotted #666;
}
.timeline-item:hover {
  background: #eee;
}
.timeline-item-body {
  font-size: 13.2pt;
  text-align: left;
}
.timeline-item-timestamp {
  color: #666;
  font-size: 10pt;
}
.timeline-item-quote-area {
  background: #efdfef;
  padding: 6px 16px;
  margin: 6px 0;
  display: flex;
  flex-direction: column;
}
.timeline-item-quote-text {
  color: #884444;
  font-size: 11pt;
}
.timeline-item-quote-timestamp {
  color: #666;
  font-size: 10pt;
}
.timeline-item-commands {
  display: flex;
  flex-direction: row;
  margin-top: 3px;
  width: 400px;
}
.timeline-item-command {
  flex: 1;
  display: flex;
  flex-direction: row;
}
.timeline-command-item {
  width: 24px;
  height: 24px;
  padding: 4px;
  border-radius: 50%;
  overflow: hidden;
  transition: .4s;
}
.timeline-command-item:hover {
  background: #b3e1ff;
}
.timeline-command-item-favs-count {
  margin: auto 4px;
}

/* 画面下部 */
.footer {
  background: #333;
  position: absolute !important;
  width: 100vw;
  height: 60px;
  left: 0px;
  bottom: 0px;
}
</style>
