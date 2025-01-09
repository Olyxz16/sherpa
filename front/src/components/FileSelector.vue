<script setup lang="ts">
import { ComboBox } from '@/components/ui/combobox'
import { useWorkstationStore } from '@/stores/workstationStore'
import type { PropType } from 'vue'
import type { Option } from '@/types/ComboBoxTypes'

const options : Option[] = [
    { value: '.env', label: '.env' },
    { value: '.env.example', label: '.env.example' }
];
const defaultValue=".env";
const searchPlaceholder="Search file";

const { setCurrentFile } = useWorkstationStore();
</script>

<template>
  <ComboBox @select="setCurrentFile" :options="options" :default-value="defaultValue" :search-placeholder="searchPlaceholder" />
</template>
