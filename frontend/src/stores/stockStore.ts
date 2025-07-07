import { defineStore } from 'pinia'
import type { Stock } from '@/types/stock'
import { getRecommendations } from '@/repository/stocksAPI/stockService'

export const useStockStore = defineStore('stock', {
    state: () => ({ recommendations: [] as Stock[] }),
    actions: {
        async fetchRecommendations() {
            this.recommendations = await getRecommendations()
        }
    }
})
