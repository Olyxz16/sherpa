<script setup lang="ts">
import { cn } from '@/lib/utils'
import { ComboBox } from '@/components/ui/combobox'
import type { Option } from '@/types/ComboBoxTypes'
import { useWorkstationStore } from '@/stores/workstationStore'
import { ref } from 'vue'

const options : Option[] = [
    { value: '.env', label: '.env' },
    { value: '.env.example', label: '.env.example' }
];
const defaultValue=".env";
const searchPlaceholder="Search file";

const { setCurrentFile } = useWorkstationStore();

const value = ref("");
function addFile() {
  const v = {
    value: value.value,
    label: value.value
  }
  if(options.includes(v)) {
     return;
  }
  options.push({
    value: value.value,
    label: value.value
  });
  value.value = "";
}
</script>

<template>
  <ComboBox @select="setCurrentFile" :options="options" :default-value="defaultValue" :search-placeholder="searchPlaceholder">
    <template #empty-msg>
      <p> No file found </p>
    </template>
    <template #end-command>
      <input type="text"
        v-model="value"
        placeholder="Add file"
        @change="addFile"
        :class="cn('flex h-11 px-2 w-full rounded-md bg-transparent py-3 text-sm outline-none placeholder:text-muted-foreground disabled:cursor-not-allowed disabled:opacity-50', $props)"
      />
    </template>
  </ComboBox>
</template>
