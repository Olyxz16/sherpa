<script setup lang="ts">
import type { HTMLAttributes } from 'vue'
import { cn } from '@/lib/utils'
import { AvatarRoot } from 'radix-vue'
import { avatarVariant, type AvatarVariants } from '.'

const props = withDefaults(defineProps<{
  class?: HTMLAttributes['class']
  shape?: AvatarVariants['shape']
}>(), {
  shape: 'circle',
})
</script>

<template>
  <AvatarRoot class="inline-flex items-center justify-center font-normal text-foreground select-none shrink-0 bg-secondary overflow-hidden h-14 w-14 text-s rounded-full">
    <slot />
  </AvatarRoot>
</template>
