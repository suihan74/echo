<template>
  <div class='hello'>
    <h1>{{ msg }}</h1>
    <h2>Essential Links</h2>
    <button @click="apiPublic">public</button>
    <button @click="apiPrivate">private</button>
    <button @click="signOut">Sign out</button>
  </div>
</template>

<script lang="ts">
import axios from 'axios'
import firebase from 'firebase'

export default {
  name: 'HelloWorld',

  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  },

  methods: {
    apiPublic: async function () {
      const res = await axios.get('http://localhost:8000/public')
      this.msg = res.data
    },

    apiPrivate: async function () {
      const res = await axios.get('http://localhost:8000/private', {
        headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
      })
      this.msg = res.data
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
  margin: 10px 0;
  padding: 10px;
}
</style>
