<template>
  <div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
    <StockCard v-for="s in stocks" :key="s.ID" :stock="s" />
  </div>
</template>

<script setup lang="ts">
import StockCard from '@/components/molecules/StockCard.vue'
import type { Stock } from '@/types/stock'
const props = defineProps<{ stocks: Stock[] }>()
</script>
