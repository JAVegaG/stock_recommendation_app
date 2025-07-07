<template>
    <DefaultLayout>
        <template #nav>
            <h1 class="text-xl font-bold">Recomendaciones</h1>
        </template>
        <StockList :stocks="stocks" />
    </DefaultLayout>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import DefaultLayout from '@/components/layouts/DefaultLayout.vue'
import StockList from '@/components/organisms/StockList.vue'
import { useStockStore } from '@/stores/stockStore'
import type { Stock } from '@/types/stock'

const store = useStockStore()
const stocks = ref<Stock[]>([])

onMounted(async () => {
    console.log('aca')
    await store.fetchRecommendations()

    console.log(store.recommendations)
    stocks.value = store.recommendations
})
</script>
