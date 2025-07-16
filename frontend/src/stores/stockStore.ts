import type { IFilterOptions } from '@/repository/stocksAPI/interfaces/settings'
import { getAll } from '@/repository/stocksAPI/stockService'
import type { Stock } from '@/types/stock'
import { defineStore } from 'pinia'

export const useStockStore = defineStore('stock', {
  state: () => ({
    isLoading: false as boolean,
    stocks: [] as Stock[],
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
    async fetchAll() {
      this.isLoading = true
      const result = await getAll(this.itemsPerPage, this.offset, this.filterOptions)
      this.isLoading = false
      this.stocks = result.Items
      this.totalItems = result.Total
    },
    setItemsPerPage(itemsPerPage: number) {
      this.itemsPerPage = Math.ceil(itemsPerPage)
      this.fetchAll()
    },
    setCurrentPage(currentPage: number) {
      this.currentPage = currentPage
      this.fetchAll()
    },
    setFilterOptions(filterOptions: IFilterOptions) {
      this.filterOptions = filterOptions
    },
  },
})
