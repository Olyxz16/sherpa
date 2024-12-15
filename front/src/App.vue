<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import HomeHeader from './components/HomeHeader.vue'
</script>

<template>
  <header>
    <div class="wrapper">
      <HomeHeader />
    </div>
  </header>

  <RouterView />

  <footer style="bottom: 1%; position: absolute; left:0; width: 100%; margin: 0; overflow: hidden;">
    <div style="text-align: center;">Made with ♥️ by <a href="https://github.com/Olyxz16"> Cédric Sazos</a> </div>
  </footer>
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

@media (min-width: 1024px) {
  /*header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }*/

}
</style>
