<script setup lang="ts">
  import { Button } from '@/components/ui/button'
  import { ref } from 'vue'
  import { storeToRefs } from 'pinia'
  import { useClipboard } from '@vueuse/core'
  import { useWorkstationStore } from '@/stores/workstationStore'

  const wsStore = useWorkstationStore();
  const { fileContent } = storeToRefs(wsStore);
  const copyToClipboard = () => {
    const source = fileContent.value;
    const { text, copy, copied, isSupported } = useClipboard({ source });
    copy(source);
  }
</script>

<template>
  <Button @click="copyToClipboard">
    <p> copy </p>
  </Button>
</template>
