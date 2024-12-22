<script setup lang="ts">
import EnvTextField from '../components/EnvTextfield.vue'
import FunctionBar from '../components/FunctionBar.vue'
import { useCookies } from '@vueuse/integrations/useCookies'
import fetchUser from '@/lib/fetchUser'

const cookies = useCookies(["session"])
const session = cookies.get("session")
if(session) {
  fetchUser(session)
}
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
