<template>
  <div id="app">
    <Slide>
      <a id="home" href="#">
        <span>Home</span>
      </a>
      <a id="file-system" href="#">
        <span>File System</span>
      </a>
      <a id="user-management" href="#users">
        <span>User</span>
      </a>
      <a id="system" href="#">
        <span>System</span>
      </a>
    </Slide>
    <img alt="Vue logo" src="./assets/logo.png">
  </div>
</template>

<script>
import { Slide } from "vue-burger-menu";

export default {
  name: 'AlphaGopher',
  components: {
    Slide
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
