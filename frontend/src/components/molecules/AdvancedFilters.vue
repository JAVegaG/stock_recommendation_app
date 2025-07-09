<template>
  <div class="relative">
    <button @click="showFilters = !showFilters" class="p-2 rounded hover:bg-gray-100">
      <FilterIcon class="w-5 h-5" />
    </button>
    <div
      v-if="showFilters"
      class="absolute right-0 mt-2 bg-white border rounded shadow-md p-4 z-10 w-72 flex flex-col justify-between gap-4"
    >
      <div class="mb-3">
        <label class="text-sm font-medium pb-2">Target To Min:</label>
        <input
          v-model.number="localFilters.TargetToMin"
          type="number"
          class="w-full border px-2 py-1 rounded"
        />
      </div>

      <div class="mb-3">
        <label class="text-sm font-medium pb-2">Target To Max:</label>
        <input
          v-model.number="localFilters.TargetToMax"
          type="number"
          class="w-full border px-2 py-1 rounded"
        />
      </div>

      <div class="mb-3">
        <label class="text-sm font-medium pb-2 block">Rating To:</label>
        <div class="flex flex-wrap gap-1">
          <button
            v-for="rating in ratings"
            :key="rating"
            @click="toggleRating(rating)"
            :class="{
              'bg-blue-500 text-white': localFilters.RatingTo === rating,
              'bg-gray-200 text-gray-700': localFilters.RatingTo !== rating,
            }"
            class="text-xs px-2 py-1 rounded"
          >
            {{ rating }}
          </button>
        </div>
      </div>

      <div class="flex justify-end gap-2 mt-2">
        <button class="text-sm text-gray-600" @click="resetFilters">Reset</button>
        <button class="text-sm font-semibold text-blue-600" @click="applyFilters">Apply</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ratings, type IFilterOptions } from '@/repository/stocksAPI/interfaces/settings'
import { FilterIcon } from 'lucide-vue-next'
import { ref, watch } from 'vue'

const props = defineProps<{
  modelValue: Omit<IFilterOptions, 'Company'>
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: Omit<IFilterOptions, 'Company'>): void
}>()

const showFilters = ref(false)
const localFilters = ref({ ...props.modelValue })

watch(
  () => props.modelValue,
  (newVal) => {
    localFilters.value = { ...newVal }
  },
)

function toggleRating(rating: string) {
  localFilters.value.RatingTo = localFilters.value.RatingTo === rating ? '' : rating
}

function resetFilters() {
  localFilters.value = {
    TargetToMin: 0,
    TargetToMax: 0,
    RatingTo: '',
  }
  emit('update:modelValue', { ...localFilters.value })
  showFilters.value = false
}

function applyFilters() {
  emit('update:modelValue', { ...localFilters.value })
  showFilters.value = false
}
</script>
