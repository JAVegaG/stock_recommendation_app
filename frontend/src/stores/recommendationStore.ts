import type { IFilterOptions } from '@/repository/stocksAPI/interfaces/settings'
import { getRecommendations } from '@/repository/stocksAPI/stockService'
import type { Stock } from '@/types/stock'
import { defineStore } from 'pinia'

export const useRecommendationStore = defineStore('recommendation', {
  state: () => ({
    isLoading: false as boolean,
    recommendations: [] as Stock[],
    currentPage: 1 as number,
    totalItems: 0 as number,
    itemsPerPage: 5 as number,
    filterOptions: {} as IFilterOptions,
  }),
  getters: {
    totalPages: (state) => Math.ceil(state.totalItems / state.itemsPerPage),
    offset: (state) => (state.currentPage - 1) * state.itemsPerPage,
  },
  actions: {
    async fetchRecommendations() {
      this.isLoading = true
      const result = await getRecommendations(this.itemsPerPage, this.offset, this.filterOptions)
      this.isLoading = false
      this.recommendations = result.Items
      this.totalItems = result.Total
    },
    setItemsPerPage(itemsPerPage: number) {
      this.itemsPerPage = Math.ceil(itemsPerPage)
      this.fetchRecommendations()
    },
    setCurrentPage(currentPage: number) {
      this.currentPage = currentPage
      this.fetchRecommendations()
    },
    setFilterOptions(filterOptions: IFilterOptions) {
      this.filterOptions = filterOptions
    },
  },
})
