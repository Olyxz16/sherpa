<script setup lang="ts">
  import LoginButton from '@/components/LoginButton.vue'
</script>

<template>
  <div class="wrapper">
    <h1 id="title"> Sherpa </h1>
    <LoginButton id="login"/>
  </div>
</template>

<style scoped>
.wrapper {
  display: flex;
  flex-direction: row;
}

#title {
  font-size: 6em;

  margin-left: 2vw;
}

#login {
  margin-left: auto;
  margin-right: 2vw;
}

</style>
