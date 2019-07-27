<template>
  <div>
    <h1>File System</h1>

    <div>
      <b-spinner 
        v-if="loading" 
        variant="primary"
        label="Spinning"
      />
    </div>
    <tree-browser
      :node="root"
      @onClick="nodeWasClicked"
    />
  </div>
</template>

<script>
import TreeBrowser from "./TreeBrowser.vue";

export default {
  name: "FileSystem",
  components: {
    "tree-browser": TreeBrowser,
  },
  data() {
    return {
      root: {},
      loading: true
    };
  },
  created() {
    this.$http
      .get("filesystem/ls/all")
      .then(function(data) {
        this.root = JSON.parse(data.body.std)[0];
        this.loading = false;
      })
      .catch(function(error) {
        console.error("failed to get files" + error.body);
        this.loading = false;
      });
  },
  methods: {
    nodeWasClicked(node) {
    }
  }
};
</script>


<style scoped>
</style>
