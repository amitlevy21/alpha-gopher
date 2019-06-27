<template>
  <div>
    <h1>File System</h1>
    <tree-browser 
      :node="root"
      @onClick="nodeWasClicked"
    />
  </div>
</template>

<script>
import TreeBrowser from './TreeBrowser.vue'
import root from './root.json'

export default {
  name: "FileSystem",
  components: {
    'tree-browser': TreeBrowser
  },
  data() {
    return {
      root,
    }
  },
  methods: {
    nodeWasClicked(node) {
      alert(node.name);
    }
  }

};
</script>


<style scoped>
</style>
