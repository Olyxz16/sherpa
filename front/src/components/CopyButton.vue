<script setup lang="ts">
  import { Button } from '@/components/ui/button'
  import { useWorkstationStore } from '@/stores/workstationStore'

  const wsstore = useWorkstationStore();
  const copyToClipboard = () => {
    const text = wsstore.fileContent();
    navigator.clipboard.writeText(text);
  }
</script>

<template>
  <Button @click="copyToClipboard">
    <p> copy </p>
  </Button>
</template>
