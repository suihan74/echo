<template>
  <div id='home'>
    <!-- 投稿エリア -->
    <div id="post-area">
      <textarea
        class="post-text"
        placeholder="いまどうしてる？"
        @keydown.prevent.ctrl.enter="post"
        v-model.trim="postText"/>
      <!-- 投稿エリア下部 -->
      <div class="post-commands">
        <span class="counter">{{ postText.length }} 文字</span>
        <button class="post-button" @click="post" title="Ctrl+Enter でも投稿できます">
          post
        </button>
      </div>
    </div>

    <div class="timeline">
      <ul class="timeline-items">
        <li class="timeline-item" v-for="post in posts" :key="post.Id">
          <span class="timeline-item-body">{{ post.Text }}</span>
          <time class="timeline-item-timestamp">{{ timeToString(post.Timestamp) }}</time>
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
import axiosBase from 'axios'
import firebase from 'firebase'
import Moment from 'moment'

// 認証情報を含んだクライアントを作成
const axios = axiosBase.create({
  baseURL: 'http://localhost:8000',
  headers: {
    'Content-Type': 'application/json',
    Authorization: `Bearer ${localStorage.getItem('jwt')}`
  },
  responseType: 'json'
})

export default {
  name: 'Home',

  data () {
    return {
      postText: '',
      posts: []
    }
  },

  methods: {
    /** 投稿する */
    post: function () {
      if (this.postText.length === 0) {
        console.error('空文字列は投稿されない')
        return
      }

      axios.post('/post', {
        Text: this.postText
      }, axios.defaults.headers).then(res => {
        // 投稿後編集内容をクリアしてタイムラインを更新する
        this.postText = ''
        this.getPosts()
      }).catch(err => {
        console.error(err)
      })
    },

    /** 最新の投稿を取得する */
    getPosts: function () {
      axios.get('posts').then(res => {
        this.posts = res.data
      }).catch(err => {
        console.error(err)
      })
    },

    signOut: function () {
      firebase.auth().signOut().then(() => {
        localStorage.removeItem('jwt')
        this.$router.push('/signin')
      })
    },

    /** タイムスタンプを表示用に加工する */
    timeToString: function (unixtime) {
      const date = new Date(unixtime * 1000)
      return Moment(date).format('YYYY-MM-DD HH:mm:ss')
    }
  },

  // ロード後にタイムラインを取得する
  mounted: function () {
    this.getPosts()
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
div#home {
  background: #fefefe;
  overflow-x: hidden;
  display: flex;
  flex-direction: column;
  margin: 0 3vw;
}

/* 投稿エリア */
div#post-area {
  flex: 1;
  background: #dfefef;
  display: flex;
  flex-direction: column;
  padding: 12px 16px;
}

/* 投稿オプションエリア */
div.post-commands {
  padding: 8px 0;
  display: flex;
  flex-direction: row;
}

/* 編集中の文字列 */
textarea.post-text {
  font-size: 14pt;
  height: 48px;
  padding: 12px;
}
button.post-button {
  margin: 0;
}
/* 文字数カウンタ */
.counter {
  font-size: 12pt;
  margin-right: auto;
}

div.timeline {
  overflow: hidden;
}
ul.timeline-items {
  margin-top: 8px;
  display: flex;
  flex-direction: column;
  float: left;
  _zoom: 1;
  overflow: hidden;
}
li.timeline-item {
  padding: 10px 16px;
  text-align: left;
  display: flex;
  flex-direction: column;
  width: 100vw;
  margin: 0;

  margin-top: -1px;
  border-top: 1px dotted #666;
}
li.timeline-item:hover {
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

/* 画面下部 */
div.footer {
  background: #333;
  position: absolute !important;
  width: 100vw;
  height: 60px;
  left: 0px;
  bottom: 0px;
}
</style>
