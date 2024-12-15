<script setup lang="ts">
import { ComboBox } from '@/components/ui/combobox'
const options = [
    { value: '.env', label: '.env' },
    { value: '.env.example', label: '.env.example' }
];
const defaultValue=".env";
const searchPlaceholder="Search file";
</script>

<template>
  <ComboBox :options="options" :default-value="defaultValue" :search-placeholder="searchPlaceholder" />
</template>
