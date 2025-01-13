<script setup lang="ts">
import { cn } from '@/lib/utils'
import { ComboBox } from '@/components/ui/combobox'
import { CommandInput } from '@/components/ui/command';
import type { Option } from '@/types/ComboBoxTypes'
import { useWorkstationStore } from '@/stores/workstationStore'
import type { PropType } from 'vue'
import { useForwardProps } from 'radix-vue'

const options : Option[] = [
    { value: '.env', label: '.env' },
    { value: '.env.example', label: '.env.example' }
];
const defaultValue=".env";
const searchPlaceholder="Search file";

const { setCurrentFile } = useWorkstationStore();
</script>

<template>
  <ComboBox @select="setCurrentFile" :options="options" :default-value="defaultValue" :search-placeholder="searchPlaceholder">
    <template #empty-msg>
      <p> No file found </p>
    </template>
    <template #end-command>
      <CommandInput class="h-9" :placeholder="searchPlaceholder" />
    </template>
  </ComboBox>
</template>
