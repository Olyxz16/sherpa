<script setup lang="ts">
  import { Button } from '@/components/ui/button'
  import { ref } from 'vue'
  import { useClipboard } from '@vueuse/core'
  import { useWorkstationStore } from '@/stores/workstationStore'

  const wsstore = useWorkstationStore();
  const copyToClipboard = () => {
    const source = wsstore.fileContent();
    const { text, copy, copied, isSupported } = useClipboard({ source });
    copy(source);
  }
</script>

<template>
  <Button @click="copyToClipboard">
    <p> copy </p>
  </Button>
</template>
