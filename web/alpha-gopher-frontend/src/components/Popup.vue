<template>
  <div class="py-4">
    <button 
      class="btn btn-primary"
      :show="show" 
      @click="show=true"
      @save="moveFile"
    >
      Move
    </button>
    <button 
      class="btn btn-primary"
      :show="show" 
      @click="show=true"
      @save="copyFile"
    >
      Copy
    </button>
    <button 
      class="btn btn-primary"
      :show="show" 
      @click="show=true"
      @save="deleteFile"
    >
      Delete
    </button>
    <stack-modal
      :show="show" 
      title="Please fill destination path below" 
      @close="show=false"
      @save="show=false"
    >
      <label>Destination</label>
      <b-form-input 
        type="text"
      /> 
    </stack-modal>
  </div>
</template>

<script>
import { BFormInput } from 'bootstrap-vue'
import StackModal from "@innologica/vue-stackable-modal";

export default {
  name: "Popup",
  components: { StackModal, BFormInput },
  props: {
    pop: {
      type: Boolean,
      default () {return false}
    }   
  },
  data() {
    return {
      show: this.pop,
      show_second: false,
      show_third: false
    };
  },
  methods: {
    moveFile(path) {
        alert(path)
    },
    copyFile(path) {
        
    },
    deleteFile(path) {
        
    }
  }
};
</script>

<style>
@import "https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css";
</style>
