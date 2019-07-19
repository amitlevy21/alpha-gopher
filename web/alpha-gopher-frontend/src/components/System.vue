<template>
  <div>
    <h1>System Tools</h1>
    <system-monitor id="htop" />
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
#htop {
  text-align: left;
}
</style>
