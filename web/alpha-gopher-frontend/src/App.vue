<template>
  <div id="app">
    <app-menu />
    <img
      alt="Vue logo"
      src="./assets/logo.png"
    >
    <div id="app-content">
      <router-view />
    </div>
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
  margin-left: 20px;
}
input {
  max-width: 200px;
  text-align: center;
}
#app-content {
  display: flex;
  justify-content: center;
  align-items: center;
}
div {
  margin: 20px
}
</style>
