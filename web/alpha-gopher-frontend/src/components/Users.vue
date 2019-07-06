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
    <input
      v-model="newUserName"
      type="text"
      required
    > 
    <button @click.prevent="addUser">
      Add User
    </button>
    <button @click.prevent="removeUser">
      Remove User
    </button>
    <div v-if="submitted">
      <h3>User Added!</h3>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Users',
  data() {
    return {
      users: [],
      newUserName: '',
      submitted: false
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
      this.$http.post('users/new/' + this.newUserName).then(function (data) {
        this.users.push(this.newUserName)
        this.submitted = true
      }).catch(function (error) {
        console.error("failed to add user error:" + error.body)
      })
    },
    removeUser() {
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
