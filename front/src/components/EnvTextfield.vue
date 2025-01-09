<script setup lang="ts">
import { ref, watch } from 'vue';
import { storeToRefs } from 'pinia';
import { Textarea } from "@/components/ui/textarea"

const placeholder = "HOST=\nPORT=";

const modelValue = ref('');
import { useWorkstationStore } from '@/stores/workstationStore'
const { updateFileContent, saveCurrentFile } = useWorkstationStore();
const wsStore = useWorkstationStore();
const { fileContent } = storeToRefs(wsStore);
watch(fileContent, async (newVal, oldVal) => {
  modelValue.value = newVal;
})
</script>

<template>
    <Textarea class="textarea" :modelValue="modelValue" @keydown.enter.shift.exact.prevent="saveCurrentFile" @keydown.s.ctrl.exact.prevent="saveCurrentFile" @update:modelValue="updateFileContent" :placeholder="placeholder"/>
</template>

<style scoped>
.textarea {
  font-size: 1.4em;
  line-height: 1em;

  width: 100%;
  height: 100%;
}
</style>
