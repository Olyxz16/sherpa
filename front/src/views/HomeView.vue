<script setup lang="ts">
import EnvTextField from '../components/EnvTextfield.vue'
import FunctionBar from '../components/FunctionBar.vue'
</script>

<template>
  <main>
    <FunctionBar />
    <EnvTextField />
  </main>
</template>

<style scoped>
main {
  width: 50vw;
  height: 60vh;

  margin-left: auto;
  margin-right: auto;
  margin-top: 5vh;
}
</style>
