<template>
  <div
    class="flex flex-col sm:flex-row items-center justify-between gap-4 p-4 bg-white rounded-xl shadow-sm border border-gray-200">
    <!-- Información de resultados -->
    <div class="text-sm text-gray-600 order-2 sm:order-1">
      <span class="font-medium">
        Showing {{ startItem }}-{{ endItem }} of {{ totalItems }}
        {{ totalItems === 1 ? 'result' : 'results' }}
      </span>
    </div>

    <!-- Controles de navegación -->
    <div class="flex items-center gap-3 order-1 sm:order-2">
      <!-- Botón Anterior -->
      <BaseButton variant="secondary" size="sm" :disabled="currentPage <= 1 || loading" :icon-left="ChevronLeft"
        @click="goToPreviousPage" class="flex-shrink-0">
        <span class="hidden sm:inline">Previous</span>
      </BaseButton>

      Current page data
      <div class="flex items-center gap-2 px-3 py-2 bg-gray-50 rounded-lg border">
        <span class="text-sm font-medium text-gray-700"> Page </span>
        <input v-model.number="pageInput" @keyup.enter="goToPage" @blur="goToPage" type="number" :min="1"
          :max="totalPagesLocal"
          class="w-12 text-center text-sm font-semibold bg-transparent border-none outline-none text-gray-800"
          :disabled="loading" />
        <span class="text-sm text-gray-500"> of {{ totalPagesLocal }} </span>
      </div>

      <!-- Botón Siguiente -->
      <BaseButton variant="secondary" size="sm" :disabled="currentPage >= totalPagesLocal || loading"
        :icon-right="ChevronRight" @click="goToNextPage" class="flex-shrink-0">
        <span class="hidden sm:inline">Next</span>
      </BaseButton>
    </div>

    <div class="flex items-center gap-2 text-sm text-gray-600 order-3">
      <span>Show:</span>
      <select v-model="itemsPerPageLocal" @change="changeItemsPerPage"
        class="px-2 py-1 border border-gray-300 rounded-md bg-white text-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        :disabled="loading">
        <option v-for="option in itemsPerPageOptions" :key="option" :value="option">
          {{ option }}
        </option>
      </select>
      <span>per page</span>
    </div>
  </div>

  <!-- Navegación rápida (solo se muestra en desktop y cuando hay muchas páginas) -->
  <div v-if="totalPagesLocal > 5 && !isMobile" class="flex justify-center mt-4">
    <div class="flex items-center gap-1 bg-white rounded-lg shadow-sm border border-gray-200 p-2">
      <!-- Primera página -->
      <button v-if="showFirstPage" @click="goToSpecificPage(1)" :disabled="loading"
        class="w-8 h-8 flex items-center justify-center text-sm font-medium rounded-md transition-colors hover:bg-gray-100 disabled:opacity-50"
        :class="currentPage === 1 ? 'bg-blue-600 text-white hover:bg-blue-700' : 'text-gray-700'">
        1
      </button>

      <!-- Puntos suspensivos izquierda -->
      <span v-if="showLeftEllipsis" class="px-2 text-gray-400">...</span>

      <!-- Páginas visibles -->
      <button v-for="page in visiblePages" :key="page" @click="goToSpecificPage(page)" :disabled="loading"
        class="w-8 h-8 flex items-center justify-center text-sm font-medium rounded-md transition-colors hover:bg-gray-100 disabled:opacity-50"
        :class="currentPage === page ? 'bg-blue-600 text-white hover:bg-blue-700' : 'text-gray-700'">
        {{ page }}
      </button>

      <!-- Puntos suspensivos derecha -->
      <span v-if="showRightEllipsis" class="px-2 text-gray-400">...</span>

      <!-- Última página -->
      <button v-if="showLastPage" @click="goToSpecificPage(totalPagesLocal)" :disabled="loading"
        class="w-8 h-8 flex items-center justify-center text-sm font-medium rounded-md transition-colors hover:bg-gray-100 disabled:opacity-50"
        :class="currentPage === totalPagesLocal ? 'bg-blue-600 text-white hover:bg-blue-700' : 'text-gray-700'
          ">
        {{ totalPagesLocal }}
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted, onUnmounted } from 'vue'
import { ChevronLeft, ChevronRight } from 'lucide-vue-next'
import BaseButton from '../atoms/BaseButton.vue'

interface Props {
  currentPage: number
  totalItems: number
  totalPages: number
  itemsPerPage?: number
  loading?: boolean
  itemsPerPageOptions?: number[]
}

const props = withDefaults(defineProps<Props>(), {
  itemsPerPage: 10,
  totalPages: 1,
  loading: false,
  itemsPerPageOptions: () => [5, 10, 20, 50],
})

const emit = defineEmits<{
  'update:currentPage': [page: number]
  'update:itemsPerPage': [items: number]
  'page-change': [page: number]
}>()

// Reactive state
const pageInput = ref(props.currentPage)
const itemsPerPageLocal = ref(props.itemsPerPage)
const totalPagesLocal = ref(props.totalPages)
const isMobile = ref(false)

// Computed properties
const startItem = computed(() => {
  if (props.totalItems === 0) return 0
  return (props.currentPage - 1) * props.itemsPerPage + 1
})

const endItem = computed(() => {
  const end = props.currentPage * props.itemsPerPage
  return Math.min(end, props.totalItems)
})

// Navegación rápida - páginas visibles
const visiblePages = computed(() => {
  const pages = []
  const current = props.currentPage
  const total = totalPagesLocal.value

  if (total <= 7) {
    // Si hay 7 páginas o menos, mostrar todas
    for (let i = 1; i <= total; i++) {
      pages.push(i)
    }
  } else {
    // Lógica para mostrar páginas alrededor de la actual
    let start = Math.max(2, current - 1)
    let end = Math.min(total - 1, current + 1)

    if (current <= 3) {
      end = 5
    } else if (current >= total - 2) {
      start = total - 4
    }

    for (let i = start; i <= end; i++) {
      pages.push(i)
    }
  }

  return pages
})

const showFirstPage = computed(() => totalPagesLocal.value > 7 && !visiblePages.value.includes(1))
const showLastPage = computed(
  () => totalPagesLocal.value > 7 && !visiblePages.value.includes(totalPagesLocal.value),
)
const showLeftEllipsis = computed(() => showFirstPage.value && visiblePages.value[0] > 2)
const showRightEllipsis = computed(
  () =>
    showLastPage.value && visiblePages.value[visiblePages.value.length - 1] < totalPagesLocal.value - 1,
)

// Methods
const goToPreviousPage = () => {
  if (props.currentPage > 1) {
    const newPage = props.currentPage - 1
    emit('update:currentPage', newPage)
    emit('page-change', newPage)
  }
}

const goToNextPage = () => {
  if (props.currentPage < totalPagesLocal.value) {
    const newPage = props.currentPage + 1
    emit('update:currentPage', newPage)
    emit('page-change', newPage)
  }
}

const goToSpecificPage = (page: number) => {
  if (page >= 1 && page <= totalPagesLocal.value && page !== props.currentPage) {
    emit('update:currentPage', page)
    emit('page-change', page)
  }
}

const goToPage = () => {
  const page = Math.max(1, Math.min(pageInput.value, totalPagesLocal.value))
  pageInput.value = page
  goToSpecificPage(page)
}

const changeItemsPerPage = () => {
  emit('update:itemsPerPage', itemsPerPageLocal.value)
  // Ajustar página actual si es necesario
  const newTotalPages = Math.ceil(props.totalItems / itemsPerPageLocal.value)
  if (props.currentPage > newTotalPages) {
    emit('update:currentPage', newTotalPages)
    emit('page-change', newTotalPages)
  }
}

const checkMobile = () => {
  isMobile.value = window.innerWidth < 768
}

// Watchers
watch(
  () => props.currentPage,
  (newPage) => {
    pageInput.value = newPage
  },
)

watch(
  () => props.itemsPerPage,
  (newItemsPerPage) => {
    totalPagesLocal.value = newItemsPerPage
  },
)

watch(
  () => totalPagesLocal.value,
  (newTotalPages) => {
    itemsPerPageLocal.value = newTotalPages
  },
)

// Lifecycle
onMounted(() => {
  checkMobile()
  window.addEventListener('resize', checkMobile)
})

onUnmounted(() => {
  window.removeEventListener('resize', checkMobile)
})
</script>
