<template>
  <div id="users">
    <h1>Users</h1>
    <div id="users-list">
      <ul>
        <li
          v-for="(user, index) in users"
          :key="index"
        >
          {{ user }}
        </li>
      </ul>
    </div>
    <b-form-input
      v-model="newUserName"
      type="text"
    /> 
    <button 
      class="btn btn-primary"
      @click="addUser"
    >
      Add User
    </button>
    <button 
      class="btn btn-primary"
      @click="removeUser"
    >
      Remove User
    </button>
    <div v-if="submitted">
      <h3>User Added!</h3>
    </div>
    <div v-if="error">
      <h3>Error when adding user</h3>
    </div>
  </div>
</template>

<script>
import { BFormInput } from 'bootstrap-vue'

export default {
  name: 'Users',
  components: { BFormInput },
  data() {
    return {
      users: [],
      newUserName: '',
      submitted: false,
      error: false
    }
  },
  created() {
    this.$http.get('users/all').then(function(data){
      this.users = data.body.users
    }).catch(function (error) {
      console.error("failed to get users" + error.body)
      
    });
  },
  methods: {
    addUser () {
      this.submitted = false
      this.error = false
      this.$http.post('users/new/' + this.newUserName).then(function (data) {
        this.users.push(this.newUserName)
        this.submitted = true
      }).catch(function (error) {
        console.error("failed to add user error:" + error.body)
        this.error = true
      })
    },
    removeUser() {
      this.submitted = false
      this.error = false
      this.$http.delete('users/' + this.newUserName).then(function (data) {
      for( var i = 0; i < this.users.length; i++){ 
        if ( this.users[i] === this.newUserName) {
          this.users.splice(i, 1); 
        }
      }
        console.log(data)
      }).catch(function (error) {
        console.log(error)
        console.error("failed to delete user error:" + error)
      })
    }
  }
}
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
