<template>
  <div>
    <label>
      add todo
      <input type="text" v-model="todoName" v-on:keyup.enter="add" placeholder="enter">
    </label>
    <ul v-if="todos.length">
      <li v-for="todo in todos" v-bind:key="todo.ID">
        <router-link v-bind:to="{ name: 'Todo', params: { id: todo.ID }}">
          {{ todo.ID }}:{{ todo.Name }}
        </router-link>
      </li>
    </ul>
    <p v-else>No todo</p>
  </div>
</template>

<script>
import http from '../module/http';

export default {
  mounted: function() {
    http.get('/todo').then(res => this.todos = res.data);
  },
  data: function() {
    return {
      todos: [],
      todoName: null,
    }
  },
  methods: {
    add: function() {
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
