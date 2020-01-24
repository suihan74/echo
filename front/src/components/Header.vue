<template>
  <div id="header">
    <h1 id="header-title">
      <router-link to="/">
        <a>echo</a>
      </router-link>
    </h1>
    <a id="sign-out" href="" @click="signOut">ログアウト</a>
  </div>
</template>

<script lang="ts">
import firebase from 'firebase'

export default {
  name: 'Header',

  props: {
    title: String
  },

  methods: {
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
a {
  color: #eee;
  text-decoration: none;
}
a#sign-out:hover {
  text-decoration: underline;
}
#header {
  background: #333;
  width: 100%;
  height: 60px;
  margin-bottom: 40px;
  position: fixed;
  left: 0;
  top: 0;
  display: flex;
  flex-direction: row;
}
#header-title {
  margin-top: auto;
  margin-bottom: auto;
  width: auto;
  padding: 0 20px;
  font-size: 22px;
  margin-right: auto;
}
#sign-out {
  margin-top: auto;
  margin-bottom: auto;
  padding-right: 12px;
  background: #00000000;
  border: 0;
  font-size: 8pt;
}
</style>
