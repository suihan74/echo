<template>
  <div class='home'>
    <h1>{{ msg }}</h1>

    <div class="post-area">
      <textarea class="post-text" placeholder="いまどうしてる？" v-model.trim="postText"/>
      <div class="post-commands">
        <span class="counter">{{ postText.length }} 文字</span>
        <button class="post-button" @click="post">post</button>
      </div>
    </div>

    <div class="footer">
      <button @click="signOut">Sign out</button>
    </div>
  </div>
</template>

<script lang="ts">
import axiosBase from 'axios'
import firebase from 'firebase'

const axios = axiosBase.create({
  baseURL: 'http://localhost:8000',
  headers: {
    'Content-Type': 'application/json'
  },
  responseType: 'json'
});

export default {
  name: 'Home',

  data () {
    return {
      msg: 'echo',
      postText: ''
    }
  },

  methods: {
    // 認証情報を付加する
    getAuthHeader: function () {
      return `Bearer ${localStorage.getItem('jwt')}`
    },

    apiPublic: async function () {
      const res = await axios.get('http://localhost:8000/public')
      this.msg = res.data
    },

    apiPrivate: async function () {
      const res = await axios.get('http://localhost:8000/private', {
        headers: { 'Authorization': this.getAuthHeader() }
      })
      this.msg = res.data
    },

    /*
     * 投稿する
     */
    post: function () {
      axios.post('http://localhost:8000/posts/post', {
        text: this.postText
      }, {
        headers: {
          'Authorization': this.getAuthHeader(),
        }
      }).then(res => {
        console.log('post: ' + res.data.Text)
      }).catch(err => {
        console.error(err)
      })
    },

    signOut: function () {
      firebase.auth().signOut().then(() => {
        localStorage.removeItem('jwt')
        this.$router.push('/signin')
      })
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

/* 各クラス */

/* 画面 */
div.home {
  overflow-x: hidden;
}

/* 投稿エリア */
div.post-area {
  background: #ccc;
  display: flex;
  flex-direction: column;
  margin: 0 3vw;
  padding: 12px 16px;
}
/* 投稿オプションエリア */
div.post-commands {
  padding: 8px 0;
  display: flex;
  flex-direction: row;
}
/* 編集中の文字列 */
.post-text {
  font-size: 20px;
  height: 60px;
  padding: 12px;
}
button.post-button {
  margin: 0;
}
/* 文字数カウンタ */
.counter {
  font-size: 16px;
  margin-right: auto;
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
