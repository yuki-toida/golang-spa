<template>
  <div>
    <p v-if="todo">{{ todo.ID }} : {{ todo.Name }}</p>
    <p v-else>No todo</p>
  </div>
</template>

<script>
import http from '../module/http'

export default {
  mounted: function() {
    const id = this.$route.params.id;
    http.get(`/todo/${id}`).then(res => this.todo = res.data);
  },
  data: function() {
    return {
      todo: null,      
    }
  },
}
</script>
