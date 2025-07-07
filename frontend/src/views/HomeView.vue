<template>
  <DefaultLayout>
    <template #nav>
      <h1 class="text-xl font-bold">Stocks </h1>
    </template>
    <template v-if="store.stocks.length > 0">
      <StockList :stocks="store.stocks" />
      <StockPagination :current-page="store.currentPage" :total-items="store.totalItems" :total-pages="store.totalPages"
        :items-per-page="store.itemsPerPage" @update:items-per-page="store.setItemsPerPage"
        @update:current-page="store.setCurrentPage" />
    </template>
    <p v-else>Data Not Found</p>
  </DefaultLayout>
</template>

<script setup lang="ts">
import DefaultLayout from '@/components/layouts/DefaultLayout.vue'
import StockPagination from '@/components/molecules/StockPagination.vue'
import StockList from '@/components/organisms/StockList.vue'
import { useStockStore } from '@/stores/stockStore'
import { onMounted } from 'vue'

const store = useStockStore()

onMounted(async () => {
  await store.fetchAll()
})
</script>
