<template>
  <div id='home'>
    <!-- 投稿エリア -->
    <div id="post-area">
      <div id="post-quote-area" v-if="quote_post">
        <div class="post-quote-body">
          <span class="quote-text">{{ quote_post.text }}</span>
          <time class="quote-timestamp">{{ timeToString(quote_post.timestamp) }}</time>
        </div>

        <button class="post-quote-close" @click="removeQuote">
          ×
        </button>
      </div>

      <textarea
        class="post-text"
        placeholder="いまどうしてる？"
        @keydown.prevent.ctrl.enter="post"
        v-model.trim="post_text"/>
      <!-- 投稿エリア下部 -->
      <div class="post-commands">
        <span class="counter">{{ post_text.length }} 文字</span>
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
              <svg class="timeline-command-item" viewBox="0 0 24 24" @click="quote(post)">
                <g id="Page-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
                    <g id="ic-message-type" fill-rule="nonzero" fill="#4A4A4A">
                        <path d="M7.34969209,20.162577 L4,21 L4,18.1947432 C2.14834206,16.5995462 1,14.4357201 1,12 C1,6.9052034 6.02423829,3 12,3 C17.9757617,3 23,6.9052034 23,12 C23,17.0947966 17.9757617,21 12,21 C10.3485759,21 8.76982106,20.7017542 7.34969209,20.162577 Z M12,19 C16.9705627,19 21,15.8659932 21,12 C21,8.13400675 16.9705627,5 12,5 C7.02943725,5 3,8.13400675 3,12 C3,15.8659932 7.02943725,19 12,19 Z M12,13 C11.4477153,13 11,12.5522847 11,12 C11,11.4477153 11.4477153,11 12,11 C12.5522847,11 13,11.4477153 13,12 C13,12.5522847 12.5522847,13 12,13 Z M8,13 C7.44771525,13 7,12.5522847 7,12 C7,11.4477153 7.44771525,11 8,11 C8.55228475,11 9,11.4477153 9,12 C9,12.5522847 8.55228475,13 8,13 Z M16,13 C15.4477153,13 15,12.5522847 15,12 C15,11.4477153 15.4477153,11 16,11 C16.5522847,11 17,11.4477153 17,12 C17,12.5522847 16.5522847,13 16,13 Z" id="Oval-1"></path>
                    </g>
                </g>
              </svg>
            </div>
            <!-- スター -->
            <div class="timeline-item-command">
              <svg class="timeline-command-item" viewBox="0 0 24 24">
                <g id="Page-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
                  <g id="ic-star" fill-rule="nonzero" fill="#4A4A4A">
                      <path d="M12.0017152,16.5646495 L12.0017152,16.5646495 L16.1590997,18.729349 C16.6992758,19.0106121 17.0545245,18.7616573 16.9526411,18.1733301 L16.1563427,13.5750922 L19.5295166,10.3186021 C19.9567491,9.90614832 19.8153364,9.49285128 19.2039409,9.40486256 L14.5694162,8.73788691 L12.4846801,4.55427191 C12.2206359,4.02439229 11.7854844,4.01899407 11.5187502,4.55427191 L9.43401412,8.73788691 L4.79948941,9.40486256 C4.1819817,9.49373092 4.04232876,9.9019464 4.4739137,10.3186021 L7.8470876,13.5750922 L7.05078926,18.1733301 C6.94993336,18.7557241 7.30418646,19.0105955 7.84433065,18.729349 L12.0017152,16.5646495 Z M13.0533494,18.4630088 L13.0533494,18.4630088 L8.89596492,20.6277083 C6.72107743,21.7601458 4.45335461,20.158052 4.8588441,17.8165462 L5.65514244,13.2183084 L6.27877839,15.1019963 L2.90560449,11.8455062 C1.16702032,10.1670626 2.06332326,7.62020285 4.47313299,7.2733961 L9.1076577,6.60642045 L7.4327014,7.800529 L9.51743752,3.61691401 C10.5944416,1.45559958 13.4144029,1.46646459 14.4859928,3.61691401 L16.5707289,7.800529 L14.8957726,6.60642045 L19.5302974,7.2733961 C21.9340082,7.61932513 22.8320473,10.1712744 21.0978259,11.8455062 L17.724652,15.1019963 L18.3482879,13.2183084 L19.1445862,17.8165462 C19.551592,20.1668077 17.2796923,21.7587604 15.1074654,20.6277083 L10.9500809,18.4630088 L13.0533494,18.4630088 Z" id="Star-1"></path>
                  </g>
                </g>
              </svg>
            </div>
            <!-- 消去 -->
            <div class="timeline-item-command">
              <svg class="timeline-command-item" v-if="post.is_yours" viewBox="0 0 24 24" @click="deletePost(post.id)">
                <g id="Page-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
                  <g id="ic-cross" fill="#4A4A4A">
                      <path d="M12,10.4834761 L7.83557664,6.31871006 C7.41207382,5.89517239 6.73224519,5.89425872 6.31350312,6.31303524 C5.89184166,6.7347314 5.89730155,7.41332336 6.31917747,7.83523399 L10.4836008,12 L6.31917747,16.164766 C5.89730155,16.5866766 5.89184166,17.2652686 6.31350312,17.6869648 C6.73224519,18.1057413 7.41207382,18.1048276 7.83557664,17.6812899 L12,13.5165239 L16.1644234,17.6812899 C16.5879262,18.1048276 17.2677548,18.1057413 17.6864969,17.6869648 C18.1081583,17.2652686 18.1026985,16.5866766 17.6808225,16.164766 L13.5163992,12 L17.6808225,7.83523399 C18.1026985,7.41332336 18.1081583,6.7347314 17.6864969,6.31303524 C17.2677548,5.89425872 16.5879262,5.89517239 16.1644234,6.31871006 L12,10.4834761 L12,10.4834761 Z" id="Combined-Shape"></path>
                  </g>
                </g>
              </svg>
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
import axiosBase from 'axios'
import firebase from 'firebase'
import Moment from 'moment'

export default {
  name: 'Home',

  data () {
    return {
      post_text: '',
      posts: [],
      quote_post: null,

      // 認証情報を含んだクライアントを作成
      axios: axiosBase.create({
        baseURL: 'http://localhost:8000',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${localStorage.getItem('jwt')}`
        },
        responseType: 'json'
      })
    }
  },

  methods: {
    /** 投稿する */
    post: function () {
      if (this.post_text.length === 0) {
        console.error('空文字列は投稿されない')
        return
      }

      const data = {
        text: this.post_text,
        quote_id: (this.quote_post ? this.quote_post.id : 0)
      }

      this.axios.post('/post', data, this.axios.defaults.headers).then(res => {
        // 投稿後編集内容をクリアしてタイムラインを更新する
        this.post_text = ''
        this.quote_post = null
        this.getPosts()
      }).catch(err => {
        console.error(err)
      })
    },

    /** (自分の)投稿を削除する */
    deletePost: function (postId) {
      this.axios.delete('/post?id=' + postId, this.axios.defaults.headers).then(res => {
        this.posts = this.posts.filter(p => p.id != postId)
        console.log('deleted post: ' + postId)
      }).catch(err => {
        console.error(err)
      })
    },

    /** 最新の投稿を取得する */
    getPosts: function () {
      this.axios.get('posts').then(res => {
        this.posts = res.data
      }).catch(err => {
        console.error(err)
      })
    },

    /** 引用対象を設定する */
    quote: function (q) {
      this.quote_post = q
      console.log('quote: ' + q.id)
    },

    /** 引用対象を除去する */
    removeQuote: function () {
      this.quote_post = null
    },

    clickPost: function (post) {
      console.log('text: ' + post.text)
      console.log('quoteId: ' + post.quote)
      console.log('quote: ' + post.quote_post)
      console.log('is_yours: ' + post.is_yours)
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
#home {
  background: #fefefe;
  overflow-x: hidden;
  display: flex;
  flex-direction: column;
  margin: 0 3vw;
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
}
.post-button {
  margin: 0;
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
  width: 100%;
}
.timeline-item-command {
  flex: 1;
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
