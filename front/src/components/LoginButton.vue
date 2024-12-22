<script setup lang="ts">
import { ref, watch } from 'vue'
import { Avatar, AvatarImage, AvatarFallback } from '@/components/ui/avatar'
import { storeToRefs } from 'pinia'
import { useUserStore } from '@/stores/userStore'
import router from '@/router'
const store = useUserStore()

const displayName = ref("Log in")
const displayAvatarUrl = ref("")

const { loggedIn, username, avatarUrl } = storeToRefs(store)
watch(loggedIn, () => {
  displayName.value = username.value
  displayAvatarUrl.value = avatarUrl.value
})

function routeLogin() {
  router.push({path: "/login"})
}
</script>

<template>
  <button @click="routeLogin">
    <div class="avatar rounded-full border border-input">
      <p id="displayname"> {{ displayName }} </p>
      <Avatar>
        <AvatarImage :src="avatarUrl" />
        <AvatarFallback>CN</AvatarFallback>
      </Avatar>
    </div>
  </button>
</template>

<style scoped>
.avatar {
  display: flex;
  flex-direction: row;
  align-items: center;

  padding: 4px;
  height: 66px;
}
#displayname {
  font-size: 1.5em;

  margin-left: 12px;
  margin-right: 12px;
}
button {
  height: 56px;
}
</style>
