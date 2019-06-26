<template>
  <Slide>
    <router-link to="/">
      <span>Home</span>
    </router-link>
    <router-link to="/file-system">
      <span>File System</span>
    </router-link>
    <router-link to="/users">
      <span>User</span>
    </router-link>
    <router-link to="/system">
      <span>System</span>
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
