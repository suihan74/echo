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
import firebase from 'firebase'

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
          localStorage.setItem('jwt', token)
          this.$router.push('/')
        })
      }, err => {
        alert(err.message)
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
