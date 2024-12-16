<script setup lang="ts">
import { ComboBox } from '@/components/ui/combobox'
import { useWorkstationStore } from '@/stores/workstationStore'

const options = [
    { value: 'go-vue-template', label: 'go-vue-template' },
    { value: 'adoutchquizz', label: 'Adoutchquizz' }
];
const placeholder="Select repository...";
const searchPlaceholder="Search repository";

const store = useWorkstationStore();
function updateCurrentFile(name: string) {
  store.setCurrentRepo(name);
}
</script>

<template>
  <ComboBox :options="options" :placeholder="placeholder" :search-placeholder="searchPlaceholder" />
</template>
