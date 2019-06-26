<template>
  <Slide>
    <a id="home" href="#">
      <span>Home</span>
    </a>
    <a id="file-system" href="#">
      <span>File System</span>
    </a>
    <router-link to="/users">
      <span>User</span>
    </router-link>
    <a id="system" href="#">
      <span>System</span>
    </a>
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
