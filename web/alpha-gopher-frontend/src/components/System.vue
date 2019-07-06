<template>
  <div>
    <h1>System Tools</h1>
    <h3>Date:</h3>
    <button @click="changeDate">
      Change System Date
    </button>
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
      date: "23/2/20"
    }
  },
  methods: {
    changeDate() {
      alert("changed!");
    },
    backup() {
      alert("changed!");
    }
  }
};
</script>

<style>
</style>
