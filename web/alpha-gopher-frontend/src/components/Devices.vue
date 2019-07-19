<template>
  <div>
    <h1>Devices</h1>
    <div>
      <b-table 
        striped
        hover
        :items="devices"
      />
    </div>
  </div>
</template>

<script>
import { BTable } from 'bootstrap-vue'

export default {

  name: 'Devices',
  components: {
    BTable
  },
  data() {
    return {
      devices: [],
      headerFields: ['Name', 'Size', 'Mount Point']
    }
  },
  created () {
    this.$http.get('/devices').then(function (data) {
      console.log(this.devices)
      this.devices = data.body.devices
    }).catch(function (error) {
      console.log(error)
      console.error("error while fetching devices")
    })
     
  }
};
</script>

<style scoped>
#devices-list {
  display: flex;
  justify-content: center;
}
.container {
  display: flex;
  justify-content: center;
}
.container ul {
  max-width: 800px;
  align-self: left;
  margin-right: auto;
  margin-left: auto;
}
.container li {
  align-self: left;
}
</style>
