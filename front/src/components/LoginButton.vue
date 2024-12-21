<script setup lang="ts">
import { Avatar, AvatarImage, AvatarFallback } from '@/components/ui/avatar'

</script>

<template>
  <button>
    <div class="avatar rounded-full border border-input">
      <p id="displayname"> Olyxz16 </p>
      <Avatar>
        <AvatarImage src="https://avatars.githubusercontent.com/u/60843442?s=96&v=4" />
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
