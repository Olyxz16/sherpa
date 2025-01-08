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

const { loggedIn, repositories } = storeToRefs(userStore);
if(loggedIn.value) {
  setOptions(repositories.value);
}
watch(loggedIn, async (newLoggedIn, oldLoggedIn) => {
  setOptions(repositories.value);
})

function setOptions(repositories: string[]) {
  options.value = repositories.map((repoName) => {
    return {value: repoName.toLowerCase() , label: repoName}
  })
}

function updateCurrentRepo(name: string) {
  wsStore.setCurrentRepo(name);
}
</script>

<template>
  <ComboBox @select="updateCurrentRepo" :options="options" :placeholder="placeholder" :search-placeholder="searchPlaceholder" />
</template>
