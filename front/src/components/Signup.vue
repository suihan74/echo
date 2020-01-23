<!--
サインアップ
-->

<template>
  <div class="signup">
    <h2>Sign up</h2>
    <input type="text" placeholder="user name" v-model="email"/>
    <input type="password" placeholder="password" v-model="password"/>
    <button @click="signUp">Register</button>
    <p>Do you have an account?
      <router-link to="/signin">Sign in</router-link>
    </p>
  </div>
</template>

<script>
import firebase from 'firebase'

export default {
  name: 'Signup',
  data () {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    signUp: function () {
      firebase.auth().createUserWithEmailAndPassword(this.email, this.password).then(res => {
        console.log('Create account: ', res.user.email)
      }).catch(error => {
        console.log(error.message)
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
.signup {
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
