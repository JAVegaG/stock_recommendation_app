<template>
  <DefaultLayout>
    <div class="flex py-12 justify-between items-center gap-4">
      <h1 class="text-xl font-bold">Stocks</h1>
      <div class="flex gap-2 items-center">
        <SearchBar v-model="company" />
        <AdvancedFilters v-model="filtersWithoutCompany" />
      </div>
    </div>

    <template v-if="store.stocks.length > 0">
      <StockList :stocks="store.stocks" />
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
import AdvancedFilters from '@/components/molecules/AdvancedFilters.vue'
import SearchBar from '@/components/molecules/SearchBar.vue'
import StockPagination from '@/components/molecules/StockPagination.vue'
import StockList from '@/components/organisms/StockList.vue'
import type { IFilterOptions } from '@/repository/stocksAPI/interfaces/settings'
import { useStockStore } from '@/stores/stockStore'
import { debounce } from 'lodash'
import { computed, onMounted, reactive, ref, watch } from 'vue'

const store = useStockStore()

const company = ref('')

const filters = reactive<Omit<IFilterOptions, 'Company'>>({
  TargetToMin: 0,
  TargetToMax: 0,
  RatingTo: '',
})

const filtersWithoutCompany = computed({
  get: () => ({
    TargetToMin: filters.TargetToMin,
    TargetToMax: filters.TargetToMax,
    RatingTo: filters.RatingTo,
  }),
  set: (value) => {
    filters.TargetToMin = value.TargetToMin
    filters.TargetToMax = value.TargetToMax
    filters.RatingTo = value.RatingTo
  },
})

const getCombinedFilters = () => ({
  Company: company.value,
  ...filtersWithoutCompany.value,
})

// Debounce solo para el search bar (company)
const debouncedFetchCompany = debounce(() => {
  store.setFilterOptions(getCombinedFilters())
  store.fetchAll()
}, 500)

watch(company, () => {
  debouncedFetchCompany()
})

// Watch para los filtros "Apply" (sin debounce)
watch(filters, () => {
  store.setFilterOptions(getCombinedFilters())
  store.fetchAll()
})

onMounted(async () => {
  await store.fetchAll()
})
</script>
