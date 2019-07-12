<template>
  <div>
    <h1>System Tools</h1>
    <button @click="backup">
      Backup System
    </button>
    <system-monitor />
  </div>
</template>

<script>
import SystemMonitor from './HTop.vue'

export default {
  components: {
    'system-monitor': SystemMonitor
  },
  data () {
    return {
      doneBackup: false
    }
  },
  methods: {
    backup() {
      this.$http.post('/backup?path=/').then(function (data) {
        this.doneBackup = true
      }).catch(function (error) {
        console.error("failed to backup " + error)
      })
    }
  }
};
</script>

<style>
</style>
