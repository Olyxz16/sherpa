<script setup lang="ts">
import GithubButton from './GithubButton.vue'
</script>

<template>
  <div class="wrapper border border-input rounded-md">
    <p> Choose source </p>
    <GithubButton />
  </div>
</template>

<style scoped>
p {
  text-align: center;
  margin-bottom: 12px;
}
.wrapper {
  padding: 12px;
}
</style>
