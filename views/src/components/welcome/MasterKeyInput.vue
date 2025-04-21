<script setup lang="ts">
import { ref } from 'vue';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button'
import { useCookies } from '@vueuse/integrations/useCookies'
import submitMasterkey from '.'

const pwd = ref('');

function update(payload: string | number) {
  pwd.value = payload.toString();
}
function submit() {
  const cookies = useCookies(["session"])
  const session = cookies.get("session")
  if(session) {
    submitMasterkey(pwd.value);
  }
}
</script>

<template>
  <div class="container">
    <Input class="masterkeyinput" @update:modelValue="update" />
    <Button @click="submit" class="masterkeybutton">
      <p> OK </p>
    </Button>
  </div>
</template>

<style scoped>
.container {
  width: 100%;

  display: flex;
  flex-direction: row;
  gap: 12px;
}
.masterkeyinput {
  width: 100%;
  max-height: 4rem;
}
.masterkeybutton {
  height: 100%;
}
</style>
