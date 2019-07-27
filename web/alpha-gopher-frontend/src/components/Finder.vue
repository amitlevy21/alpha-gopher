<template>
  <div>
    <h1>File Finder</h1>
    <b-form-input
      v-model="filePath"
      placeholder="enter file name"
      type="text"
    />
    <b-form-input
      v-model="date"
      type="date"
    />
    <button 
      class="btn btn-primary"
      @click="find"
    >
      Find
    </button>
    <div>
      <b-list-group>
        <b-list-group-item
          v-for="(res, index) in results"
          :key="index"
        >
          {{ res }} 
        </b-list-group-item>
      </b-list-group>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      filePath: '',
      results: [],
      date: ''
    }
  },
  methods: {
    find () {
      console.log(this.date)
      this.$http.get(`find?path=/&opts=-name&opts=${this.filePath}&opts=-newermt&opts=${this.date}`).then(data => {
        this.results = data.body.matches
      })
    }
  }
}
</script>

<style>

</style>
