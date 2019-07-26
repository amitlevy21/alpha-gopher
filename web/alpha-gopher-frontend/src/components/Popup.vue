<template>
  <div class="py-4">
    <button 
      class="btn btn-primary"
      :show="show" 
      @click="show=true;methodToUse=moveFile"
    >
      Move
    </button>
    <button 
      class="btn btn-primary"
      :show="show" 
      @click="show=true;methodToUse=copyFile"
    >
      Copy
    </button>
    <button 
      class="btn btn-primary"
      :show="show" 
      @click="show=true;methodToUse=deleteFile"
    >
      Delete
    </button>
    <stack-modal
      :show="show" 
      title="Please fill destination path below" 
      @close="show=false"
      @save="show=false;methodToUse()"
    >
      <label>Destination</label>
      <b-form-input 
        v-model="text"
        placeholder="enter destination path"
        type="text"
      /> 
    </stack-modal>
    <stack-modal
      :show="show_confirm" 
      title="Done!" 
      @close="show_confirm=false"
    />
    <stack-modal
      :show="show_error" 
      title="An error occurod during the action" 
      @close="show_error=false"
    />
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
    },
    node: {
      type: Object,
      default () {return {}}
    },    
  },
  data() {
    return {
      show: this.pop,
      show_confirm: false,
      show_error: false,
      text: '',
      methodToUse: ''
    };
  },
  methods: {
    moveFile() {
      this.$http.post(`filesystem/mv?src=${this.node.name}&dst=${this.text}`).then(data => {
        this.show_confirm = true
      }).catch(err => {
        this.show_error = true
      })
    },
    copyFile() {
      this.$http.post(`filesystem/cp?src=${this.node.name}&dst=${this.text}`).then(data => {
        this.show_confirm = true
      }).catch(err => {
        this.show_error = true
      })
    },
    deleteFile() {
      this.$http.post(`filesystem/rm?dst=${this.node.name}`).then(data => {
        this.show_confirm = true
      }).catch(err => {
        console.log(err)
        this.show_error = true
      })
    }
  }
};
</script>

<style>
@import "https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css";
</style>
