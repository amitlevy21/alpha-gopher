<template>
  <div>
    <h1>System Tools</h1>
    <div id="div-htop">
      <system-monitor id="htop" />
    </div>
  </div>
</template>

<script>
import SystemMonitor from './HTop.vue'

export default {
  components: {
    'system-monitor': SystemMonitor,
  },
};
</script>

<style scoped>
#div-htop {
  text-align: center;
}
#htop {
  display: inline-block
}
</style>
