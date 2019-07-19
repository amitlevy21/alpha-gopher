<template>
  <sys-monitor id="htop" />
</template>

<script>
import SysMonitor from '../../public/htop.html'

var body = document.getElementsByTagName('body')[0];
console.log(body)
export default {
    components: {
        'sys-monitor': SysMonitor
    }
}
</script>

<style scoped>
#htop {
  background: black;
}
</style>
