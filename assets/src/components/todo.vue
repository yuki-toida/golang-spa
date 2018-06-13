<template>
  <div>
    <div v-if="todo">
      <div>ID : {{ todo.ID }}</div>
      <div>Name : <input type="text" v-model="todoName"></div>
      <div>CreatedAt : {{ todo.CreatedAt }}</div>
      <div>UpdatedAt : {{ todo.UpdatedAt }}</div>
      <div>
        <button v-on:click="update">更新</button>
        <button v-on:click="remove">削除</button>
      </div>
    </div>
    <p v-if="todo"></p>
  </div>
</template>

<script>
import http from '../module/http'

export default {
  created: function() {
    const id = this.$route.params.id;
    http.get(`/todo/${id}`).then(res => {
      this.todo = res.data;
      this.todoName = this.todo.Name;
    });
  },
  data: function() {
    return {
      todo: null,
      todoName: null
    }
  },
  methods: {
    update: function() {
      if (this.todoName) {
        http.put(`/todo/${this.todo.ID}`, {name: this.todoName}).then(res => {
          if (res.data) {
            console.log(res.data);
            this.todo = res.data;
          }
        })
      }
    },
    remove: function() {
      http.delete(`/todo/${this.todo.ID}`).then(res => {
        if (res.data) {
          this.$router.push({name: 'Todos'});
        }
      })
    }
  }
}
</script>
