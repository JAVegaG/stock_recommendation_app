<template>
    <div class="border rounded p-4 shadow">
        <h2 class="text-lg font-bold">{{ stock.ticker }}</h2>
        <p class="text-sm text-gray-600">{{ stock.company }}</p>
        <BaseButton>Ver más</BaseButton>
    </div>
</template>

<script setup lang="ts">
import BaseButton from '@/components/atoms/BaseButton.vue'
import type { Stock } from '@/types/stock'
const props = defineProps<{ stock: Stock }>()
</script>
