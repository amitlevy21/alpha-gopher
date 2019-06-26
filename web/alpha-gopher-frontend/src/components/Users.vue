<template>
  <div id="users">
    <h1>Users</h1>
    <div id="users-list">
      <ul>
        <li v-for="user in users">
          {{ user }}
        </li>
      </ul>
    </div>
    <input type="text" v-model="newUserName" required /> 
    <button @:click="addUser">
      Add User
    </button>
  </div>
</template>

<script>
export default {
  name: 'Users',
  data() {
    return {
      users: [],
      newUserName: ''
    }
  },
  created() {
    this.$http.get('users/all').then(function(data){
            this.users = data.body
        }).catch(function (error) {
          console.log(error);
          
          console.log("failed to get users" + error.body)
        });
  },
  methods: {
    addUser: function () {
      this.$http.post('users/new/' + this.newUserName).then(function (data) {

      }).catch(function (error) {
        console.log("failed to add user error:" + error)
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
