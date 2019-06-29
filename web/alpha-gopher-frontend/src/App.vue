<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <app-menu />
    <router-view />
  </div>
</template>

<script>
import Menu from "./components/Menu.vue"

export default {
  name: 'AlphaGopher',
  components: {
    'app-menu': Menu
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
