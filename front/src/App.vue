<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  data () {
    return {
      title: 'echo'
    }
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css?family=M+PLUS+Rounded+1c&display=swap');

#app {
  font-family: 'M PLUS Rounded 1c', sans-serif;
/*  font-family: 'Avenir', Helvetica, Arial, sans-serif;*/
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin: 0px;
  padding-top: 52px;
}
#router-view {
  overflow-y: auto;
  overflow-x: hidden;
}
</style>
