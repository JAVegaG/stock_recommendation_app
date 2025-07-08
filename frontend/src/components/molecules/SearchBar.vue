<template>
  <div class="relative w-full max-w-sm">
    <input
      type="text"
      :value="modelValue"
      @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
      placeholder="Type a company's name"
      class="w-full border border-gray-300 rounded px-4 py-2 pr-10 focus:outline-none focus:ring-2 focus:ring-blue-500"
    />
    <button
      v-if="modelValue"
      @click="$emit('update:modelValue', '')"
      class="absolute right-2 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
    >
      ✕
    </button>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  modelValue?: string
}>()
defineEmits<{
  (event: 'update:modelValue', value: string): void
}>()
</script>
