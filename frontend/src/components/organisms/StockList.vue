<template>
  <div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
    <StockCard v-if="!isLoading" v-for="s in stocks" :key="s.ID" :stock="s" />
    <StockCardSkeleton v-else v-for="s in new Array(itemsPerPage)" />
  </div>
</template>

<script setup lang="ts">
import StockCard from '@/components/molecules/StockCard.vue'
import StockCardSkeleton from '@/components/molecules/StockCardSkeleton.vue'
import type { Stock } from '@/types/stock'

const props = defineProps<{ stocks: Stock[]; isLoading: boolean; itemsPerPage: number }>()
</script>
