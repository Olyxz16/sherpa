<script setup lang="ts">
import { ref, watch } from 'vue'
import { storeToRefs } from 'pinia'
import { ComboBox } from '@/components/ui/combobox'
import type { Option } from '@/types/ComboBoxTypes'
import { useWorkstationStore } from '@/stores/workstationStore'
const wsStore = useWorkstationStore();
import { useUserStore } from '@/stores/userStore'
const userStore = useUserStore()

const placeholder="Select repository...";
const searchPlaceholder="Search repository";
const options = ref<Option[]>([]);

const { loggedIn, repositories } = storeToRefs(userStore)
watch(loggedIn, () => {
  options.value = repositories.value.map((repoName) => {
    return {value: repoName.toLowerCase() , label: repoName}
  })
})

function updateCurrentFile(name: string) {
  wsStore.setCurrentRepo(name);
}
</script>

<template>
  <ComboBox :options="options" :placeholder="placeholder" :search-placeholder="searchPlaceholder" />
</template>
