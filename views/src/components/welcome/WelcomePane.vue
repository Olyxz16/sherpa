<script setup lang="ts">
  import MasterKeyInput from './MasterKeyInput.vue'
</script>

<template>
  <div id="welcome-pane" class="wrapper border border-input rounded-md">
    <p class="text-2xl"> Welcome to Sherpa ! </p>
    <p class="text-base"> Please enter a password </p>
    <p class="text-sm pb-4"> It will be used to encrypt your files.</p>
    <MasterKeyInput />
  </div>
</template>

<style scoped>
#welcome-pane {
  margin: 0;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);

  width: fit-content;
  height: fit-content;

  padding: 24px;

  display: flex;
  flex-direction: column;
  align-items: center;
}
</style
