<script setup lang="ts">
import { ref } from 'vue';
import { Textarea } from "@/components/ui/textarea"
import { useWorkstationStore } from '@/stores/workstationStore'

const placeholder = "HOST=\nPORT=";

const { updateFileContent, saveCurrentFile } = useWorkstationStore();
</script>

<template>
    <Textarea class="textarea" @keydown.enter.shift.exact.prevent="saveCurrentFile" @keydown.s.ctrl.exact.prevent="saveCurrentFile" @update:modelValue="updateFileContent" :placeholder="placeholder"/>
</template>

<style scoped>
.textarea {
  font-size: 1.4em;
  line-height: 1em;

  width: 100%;
  height: 100%;
}
</style>
