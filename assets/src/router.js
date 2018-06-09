import Vue from 'vue'
import Router from 'vue-router'
import Home from './components/home.vue'
import Todos from './components/todos.vue'
import Todo from './components/todo.vue'

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/todo',
      name: 'Todos',
      component: Todos
    },
    {
      path: '/todo/:id',
      name: 'Todo',
      component: Todo
    },
  ]
});
