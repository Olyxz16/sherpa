<script setup lang="ts">
import router from '@/router'

const props = defineProps(['path'])

function routeTo() {
  router.push({path: props.path})
}
</script>

<template>
  <div @click="routeTo" class="view-blocker">
    <slot></slot>
  </div>
</template>

<style scoped>
  .view-blocker {
    background-color: rgb(0, 0, 0);
    opacity: 80%;

    position: absolute;
    left: 0px;
    bottom: 0px;
    z-index: 9999;

    width: 100vw;
    height: 100vh;
  }
</style>
