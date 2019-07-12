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

export default {
  name: "FileSystem",
  components: {
    'tree-browser': TreeBrowser
  },
  data () {
    return {
      root: {}
    }
  },
  created() {
    this.$http.get('filesystem/ls/all').then(function (data) {
      console.log(data)
      this.root = JSON.parse(data.body.std)[0]
    }).catch(function (error) {
      console.error("failed to get files" + error.body)
    });
  },
  methods: {
    nodeWasClicked(node) {
      alert(node.name);
    }
  },


};
</script>


<style scoped>
</style>
