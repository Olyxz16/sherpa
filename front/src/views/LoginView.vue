<script setup lang="ts">
import HomeView from './HomeView.vue';
import LoginForm from '@/components/login/LoginForm.vue';
import ViewBlocker from '@/components/ViewBlocker.vue'
</script>

<template>
  <HomeView />
  <ViewBlocker path="/">
    <LoginForm id="login-form" />
  </ViewBlocker>
</template>

<style scoped>
  #login-form {
    margin: 0;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }

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
