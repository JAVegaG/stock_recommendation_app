import { defineStore } from 'pinia'
import type { Stock } from '@/types/stock'
import { getAll, getRecommendations } from '@/repository/stocksAPI/stockService'

export const useRecommendationStore = defineStore('recommendation', {
  state: () => ({
    recommendations: [] as Stock[],
    currentPage: 1 as number,
    totalItems: 0 as number,
    itemsPerPage: 5 as number,
  }),
  getters: {
    totalPages: (state) => Math.ceil(state.totalItems / state.itemsPerPage),
  },
  actions: {
    async fetchRecommendations(offset: number = 0) {
      const result = await getRecommendations(this.itemsPerPage, offset)
      this.recommendations = result.Items
      this.totalItems = result.Total
    },
    setItemsPerPage(itemsPerPage: number) {
      this.itemsPerPage = Math.ceil(itemsPerPage)
      const offset = (this.currentPage - 1) * this.itemsPerPage
      this.fetchRecommendations(offset)
    },
    setCurrentPage(currentPage: number) {
      this.currentPage = currentPage
      const offset = (this.currentPage - 1) * this.itemsPerPage
      this.fetchRecommendations(offset)
    },
  },
})
