<template>
  <RouterHolder />
</template>

<script>
import RouterHolder from './view/routerHolder.vue'

export default {
  name: 'App',
  components: {
    RouterHolder
  }
}
</script>

<style>
html, body {
  padding: 0px;
  margin: 0px;
  width: 100%;
  min-height: 100vh;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
