<template>
  <div id="header">
    <h1 class="header-title">
      <router-link to="/">
        <a>echo</a>
      </router-link>
    </h1>
  </div>
</template>

<script lang="ts">
export default {
  name: 'Header',

  props: {
    title: String
  }
}
</script>

<style scoped>
a {
  color: #eee;
  text-decoration: none;
}
#header {
  background: #333;
  width: 100vw;
  height: 60px;
  margin-bottom: 40px;
  position: fixed;
  left: 0;
  top: 0;
}
.header-title {
  vertical-align: middle;
  float: left;
  padding: 0 20px;
  font-size: 22px;
  color: #eee;
}
</style>
