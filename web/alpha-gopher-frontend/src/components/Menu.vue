<template>
  <Slide id="the-slide">
    <router-link to="/">
      <span>Home</span>
    </router-link>
    <router-link to="/filesystem">
      <span>File System</span>
    </router-link>
    <router-link to="/users">
      <span>User</span>
    </router-link>
    <router-link to="/system">
      <span>System</span>
    </router-link>
    <router-link to="/backup">
      <span>Backup</span>
    </router-link>
    <router-link to="/terminal">
      <span>Terminal</span>
    </router-link>
    <router-link to="/help">
      <span>Help</span>
    </router-link>
  </Slide>
</template>

<script>
import { Slide } from "vue-burger-menu";

export default {
  name: "Menu",
  components: {
    Slide
  }
}
</script>


<style scoped>
#the-slide {
      position: fixed; /* Stay in place */
      background-color: rgb(63, 63, 65); /* Black*/
      padding-top: 24px; /* Place content 60px from the top */
      z-index: 1000; /* Stay on top */
      top: 100;
      left: 0;
    }
</style>
