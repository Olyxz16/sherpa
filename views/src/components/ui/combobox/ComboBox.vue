<script setup lang="ts">
import { Button } from '@/components/ui/button'
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from '@/components/ui/command'

import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover'
import { cn } from '@/lib/utils'
import { Check, ChevronsUpDown } from 'lucide-vue-next'
import { ref, type PropType } from 'vue'
import type { Option } from '@/types/ComboBoxTypes'

const props = defineProps({
  options: {
    type: Array as PropType<Option[]>,
    required: true
  },
  placeholder: String,
  searchPlaceholder: String,
  defaultValue: String
});

const emit = defineEmits(['select']);
function emitSelect(payload: string) {
  emit('select', payload);
}

const open = ref(false)
const value = ref(props.defaultValue)
</script>

<template>
  <Popover v-model:open="open">
    <PopoverTrigger as-child>
      <Button
        variant="outline"
        role="combobox"
        :aria-expanded="open"
        class="w-[200px] justify-between"
      >
        {{ value
          ? options.find((opt) => opt.value === value)?.label
        : placeholder }}
        <ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
      </Button>
    </PopoverTrigger>
    <PopoverContent class="w-[200px] p-0">
      <Command>
      <CommandInput class="h-9" :placeholder="searchPlaceholder" />
        <CommandEmpty>
          <slot name="empty-msg"> </slot>
        </CommandEmpty>
        <CommandList>
          <CommandGroup>
            <CommandItem
              v-for="opt in options"
              :key="opt.value"
              :value="opt.value"
              @select="(ev) => {
                if (typeof ev.detail.value === 'string') {
                  value = ev.detail.value
                }
                value ??= ''
                open = false
                emitSelect(value);
              }"
            >
              {{ opt.label }}
              <Check
                :class="cn(
                  'ml-auto h-4 w-4',
                  value === opt.value ? 'opacity-100' : 'opacity-0',
                )"
              />
            </CommandItem>
          </CommandGroup>
        </CommandList>
      <slot name="end-command"></slot>
      </Command>
    </PopoverContent>
  </Popover>
</template>
