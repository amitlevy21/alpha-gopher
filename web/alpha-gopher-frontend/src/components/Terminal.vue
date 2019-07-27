<template>
  <div>
    <h1>Terminal</h1>
    <app-terminal
      intro="Welcome to the Alpha Gopher terminal, you are running as root inside a container"
      console-sign="$"
      allow-arbitrary
      height="500px"
      @command="onCliCommand"
    />
  </div>
</template>

<script>
import VueTerminal from "./VueTerminal";
import $ from 'jquery'

export default {
  components: {
    'app-terminal': VueTerminal
  },
  data() {
    return {
      ptty: null  
    }
  },
  methods: {
      onCliCommand (data, resolve, reject) {
        console.log(data)
        this.$http.post('terminal/exec/?command=' + data.text).then(function (res) {
          console.log(res.body.std)
          resolve(res.body.std)
        }).catch(function (error) {
          console.log(error)
          resolve(error.body.std)
          console.error("failed to execute command:" + error.body)
        })
      }
  }
};
</script>

<style>
</style>
