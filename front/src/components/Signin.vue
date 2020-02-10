<!--
サインイン
-->

<template>
  <div class="signin">
    <h2>Sign in</h2>
    <input type="text" placeholder="user name" v-model="email"/>
    <input type="password" placeholder="password" v-model="password"/>
    <button @click="signIn">Signin</button>
    <p>Do you have no accounts?
      <router-link to="/signup">Sign up</router-link>
    </p>
  </div>
</template>

<script>
import * as firebase from 'firebase/app'
import 'firebase/auth'

import axios from 'axios'
import { Service } from 'axios-middleware'

// axiosに処理を挟み込む
const service = new Service(axios)
service.register({
  onRequest (config) {
    config.baseURL = 'http://localhost:8000'
    if (config.auth) {
      config.headers.Authorization = `Bearer ${config.auth}`
      delete config.auth
    }
    return config
  }
})

export default {
  name: 'Signin',
  data () {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    signIn: function () {
      firebase.auth().signInWithEmailAndPassword(this.email, this.password).then(res => {
        res.user.getIdToken().then(token => {
          this.signInServer(token)
        })
      }, err => {
        alert(err.message)
      })
    },

    signInServer: function (token) {
      axios.get('auth', { auth: token }).then(res => {
        localStorage.setItem('jwt', token)
        this.$router.push('/home')
      }, err => {
        alert(err.message)
        localStorage.removeItem('jwt')
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
.signin {
  margin-top: 20px;
  display: flex;
  flex-flow: column nowrap;
  justify-content: center;
  align-items: center
}
input {
  margin: 10px 0;
  padding: 10px;
}
button {
  margin: 10px 0;
  padding: 10px;
}
</style>
