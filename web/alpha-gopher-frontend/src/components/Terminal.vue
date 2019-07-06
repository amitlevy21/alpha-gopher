<template>
  <div>
    <h1>Terminal</h1>
    <app-terminal
      :intro="intro"
      console-sign="$"
      allow-arbitrary
      height="500px"
      @command="onCliCommand"
    />
  </div>
</template>

<script>
import VueTerminal from "vue-terminal-ui";

export default {
  components: {
    'app-terminal': VueTerminal
  },
  methods: {
      onCliCommand (data, resolve, _) {
        console.log(data.text)
        resolve('')
      }
  }
};
</script>

<style>
</style>
