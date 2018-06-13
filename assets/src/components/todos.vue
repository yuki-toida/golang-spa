<template>
  <div>
    <label>
      <input type="text" v-model="todoName" placeholder="todo">
      <button v-on:click="create">追加</button>
    </label>
    <table v-if="todos.length" border="1">
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>CreatedAt</th>
          <th>UpdatedAt</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="todo in todos" v-bind:key="todo.ID">
          <th>
            <router-link v-bind:to="{ name: 'Todo', params: { id: todo.ID }}">
              {{ todo.ID }}
            </router-link>
          </th>
          <th>{{ todo.Name }}</th>
          <th>{{ todo.CreatedAt }}</th>
          <th>{{ todo.UpdatedAt }}</th>
        </tr>
      </tbody>
    </table>
    <p v-else>No todos</p>
  </div>
</template>

<script>
import http from '../module/http';

export default {
  created: function() {
    http.get('/todo').then(res => this.todos = res.data);
  },
  data: function() {
    return {
      todos: [],
      todoName: null,
    }
  },
  methods: {
    create: function() {
      if (this.todoName) {
        http.post('/todo', {name: this.todoName}).then(res => {
          console.log(res);
          this.todos = res.data;
          this.todoName = null;
        })
      } else {
        alert('empty');
      }
    },
  }
}
</script>
