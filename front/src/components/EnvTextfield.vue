<script setup lang="ts">
import { ref } from 'vue';
import { Textarea } from "@/components/ui/textarea"
import { useWorkstationStore } from '@/stores/workstationStore'

const placeholder = "HOST=\nPORT=";

const wsstore = useWorkstationStore();
const updateStore = (payload : string | number) => {
  wsstore.updateFileContent(payload)
}
</script>

<template>
    <Textarea class="textarea" @update:modelValue="updateStore" :placeholder="placeholder"/>
</template>

<style scoped>
.textarea {
  font-size: 1.4em;
  line-height: 1em;

  width: 100%;
  height: 100%;
}
</style>
