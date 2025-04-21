<script setup lang="ts">
import EnvTextField from '../components/EnvTextfield.vue'
import FunctionBar from '../components/FunctionBar.vue'
import { tryLoadUserData } from '@/stores/userStore'

tryLoadUserData();
</script>

<template>
  <main>
    <div id="home">
      <FunctionBar />
      <EnvTextField />
    </div>
  </main>
</template>

<style scoped>
#home {
  width: 50vw;
  height: 60vh;

  margin-left: auto;
  margin-right: auto;
  margin-top: 5vh;
}
</style>
