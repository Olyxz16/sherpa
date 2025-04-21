<script setup lang="ts">
import RepositorySelector from '@/components/RepositorySelector.vue'
import FileSelector from '@/components/FileSelector.vue'
import CopyButton from '@/components/CopyButton.vue'
import DownloadButton from '@/components/DownloadButton.vue'
</script>

<template>
  <div class="wrapper">
    <RepositorySelector />
    <FileSelector />
    <CopyButton class="copy"/>
    <DownloadButton />
  </div>
</template>


<style scoped>
.wrapper {
  display: flex;
  flex-direction: row;
  gap: 5px;

  margin-bottom: 5px;
}

.copy {
  margin-left: auto;
}
</style>
