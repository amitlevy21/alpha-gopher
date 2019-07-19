<template>
  <div>
    <h1>Backup System</h1>
    <loading-button
      :loading="backInProgress"
      class="btn btn-primary"
      :show="show"
      @click.native="backup"
    >
      Backup System
    </loading-button>
  </div>
</template>

<script>
import VueLoadingButton from "vue-loading-button";
export default {
  name: "Backup",
  components: {
    "loading-button": VueLoadingButton
  },
  data() {
    return {
      doneBackup: false,
      backInProgress: false
    };
  },
  methods: {
    backup() {
      this.backInProgress = true;
      this.$http
        .post("/backup?path=/")
        .then(function(data) {
          this.doneBackup = true;
          this.backInProgress = false;
        })
        .catch(function(error) {
          console.error("failed to backup " + error);
          this.backInProgress = false;
        });
    }
  }
};
</script>

<style>
</style>
