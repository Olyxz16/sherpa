<script setup lang="ts">
import { ComboBox } from '@/components/ui/combobox'
import { useWorkstationStore } from '@/stores/workstationStore';

const options = [
    { value: '.env', label: '.env' },
    { value: '.env.example', label: '.env.example' }
];
const defaultValue=".env";
const searchPlaceholder="Search file";

const store = useWorkstationStore();
function updateCurrentFile(name: string) {
  store.setCurrentFile(name);
}
</script>

<template>
  <ComboBox @select="updateCurrentFile" :options="options" :default-value="defaultValue" :search-placeholder="searchPlaceholder" />
</template>
