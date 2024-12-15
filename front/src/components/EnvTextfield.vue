<script setup lang="ts">
import { Textarea } from "@/components/ui/textarea"
const placeholder = "HOST=\nPORT=";
</script>

<template>
    <Textarea class="textarea" :placeholder="placeholder"/>
</template>

<style scoped>
.textarea {
  font-size: 1.4em;
  line-height: 1em;

  width: 100%;
  height: 100%;
}
</style>
