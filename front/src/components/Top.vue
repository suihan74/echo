<template>
  <div class='top'>
    <h1>echo</h1>
    <h2>完全匿名SNS</h2>
    <button @click="signIn">ログイン</button>
    <button @click="signUp">新規登録</button>
  </div>
</template>

<script lang="ts">
export default {
  name: 'Top',
  methods: {
    signIn: async function () {
      this.$router.push('/signin')
    },
    signUp: async function () {
      this.$router.push('/signup')
    }
  }
}
</script>

<style scoped>
h1, h2 {
  font-weight: normal;
  align-content: center;
}
a {
  color: #42b983;
}
button {
  margin: 10px 0;
  padding: 10px;
}
</style>
