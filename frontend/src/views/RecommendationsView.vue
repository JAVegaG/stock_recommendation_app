<template>
  <DefaultLayout>
    <template #nav>
      <h1 class="text-xl font-bold">Recommendations</h1>
    </template>
    <template v-if="store.recommendations.length > 0">
      <StockList :stocks="store.recommendations" />
      <StockPagination
        :current-page="store.currentPage"
        :total-items="store.totalItems"
        :total-pages="store.totalPages"
        :items-per-page="store.itemsPerPage"
        @update:items-per-page="store.setItemsPerPage"
        @update:current-page="store.setCurrentPage"
      />
    </template>
    <p v-else>Data Not Found</p>
  </DefaultLayout>
</template>

<script setup lang="ts">
import DefaultLayout from '@/components/layouts/DefaultLayout.vue'
import StockPagination from '@/components/molecules/StockPagination.vue'
import StockList from '@/components/organisms/StockList.vue'
import { useRecommendationStore } from '@/stores/recommendationStore'
import { onMounted } from 'vue'

const store = useRecommendationStore()

onMounted(async () => {
  await store.fetchRecommendations()
})
</script>
