<script setup lang="ts">
import { ComboBox } from '@/components/ui/combobox'
const options = [
    { value: 'go-vue-template', label: 'go-vue-template' },
    { value: 'adoutchquizz', label: 'Adoutchquizz' }
];
const placeholder="Select repository...";
const searchPlaceholder="Search repository";
</script>

<template>
  <ComboBox :options="options" :placeholder="placeholder" :search-placeholder="searchPlaceholder" />
</template>
