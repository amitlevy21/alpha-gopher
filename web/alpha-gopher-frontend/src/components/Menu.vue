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
    <router-link to="/devices">
      <span>Devices</span>
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

</style>
