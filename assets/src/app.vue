<template>
  <div>
    <h1>SPA for vue-router</h1>
    <ul>
      <li><router-link v-bind:to="{ name: 'Home' }">Home</router-link></li>
      <li><router-link v-bind:to="{ name: 'Todos' }">Todos</router-link></li>
    </ul>
    <hr>
    <router-view></router-view>
  </div>
</template>
