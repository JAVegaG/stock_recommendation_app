<template>
  <div class="max-w-sm mx-auto">
    <!-- Basic Card View -->
    <div v-if="!showDetails"
      class="bg-gradient-to-br from-white to-gray-50 rounded-2xl shadow-lg hover:shadow-xl transition-all duration-300 border border-gray-200 overflow-hidden h-[27rem] flex flex-col gap-2">
      <!-- Header with Ticker -->
      <div class="bg-gradient-to-r from-blue-600 to-blue-700 p-4 text-white">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-2xl font-bold">{{ stock.Ticker }}</h3>
            <p class="text-blue-100 text-sm opacity-90">{{ stock.Company }}</p>
          </div>
          <div class="bg-white/20 rounded-full p-2">
            <TrendingUp class="w-6 h-6" />
          </div>
        </div>
      </div>

      <!-- Price Range -->
      <div class="p-6 space-y-4 flex flex-col gap-4 h-full justify-between">
        <div class="grid grid-cols-2 gap-4">
          <div class="bg-green-50 border border-green-200 rounded-xl p-4 text-center">
            <div class="text-green-600 text-xs font-semibold uppercase tracking-wide mb-1">
              Target Min
            </div>
            <div class="text-green-800 text-xl font-bold">${{ stock.TargetFrom.toFixed(2) }}</div>
          </div>
          <div class="bg-blue-50 border border-blue-200 rounded-xl p-4 text-center">
            <div class="text-blue-600 text-xs font-semibold uppercase tracking-wide mb-1">
              Target Max
            </div>
            <div class="text-blue-800 text-xl font-bold">${{ stock.TargetTo.toFixed(2) }}</div>
          </div>
        </div>

        <!-- Action Badge -->
        <template v-if="!!getActionIcon(stock.RatingTo)">
          <div class="flex justify-center">
            <span :class="[
              'inline-flex items-center px-4 py-2 rounded-full text-sm font-semibold',
              getActionColor(stock.RatingTo),
            ]">
              <component :is="getActionIcon(stock.RatingTo)" class="w-4 h-4 mr-2" />
              {{ stock.RatingTo }}
            </span>
          </div>
        </template>

        <!-- Update Info -->
        <div class="flex items-center justify-between text-sm text-gray-600 bg-gray-50 rounded-lg p-3">
          <div class="flex items-center w-32">
            <Clock class="w-4 h-4 mr-2" />
            <span>Updated {{ formatTimeAgo(stock.Time) }}</span>
          </div>
          <p class="text-xs text-gray-500 w-32">{{ stock.Brokerage }}</p>
        </div>
      </div>

      <!-- Ver más button -->
      <div class="px-6 pb-6" style="margin-top: auto">
        <BaseButton @click="showDetails = true" class="w-full group">
          <span class="flex items-center justify-center">
            View details
            <ChevronRight class="w-4 h-4 ml-2 group-hover:translate-x-1 transition-transform" />
          </span>
        </BaseButton>
      </div>
    </div>

    <!-- Detailed Card View -->
    <div v-else
      class="bg-gradient-to-br from-white to-gray-50 rounded-2xl shadow-xl border border-gray-200 overflow-hidden w-[22rem]">
      <!-- Header -->
      <div class="bg-gradient-to-r from-slate-800 to-slate-900 p-6 text-white">
        <div class="flex items-start justify-between">
          <div>
            <h2 class="text-3xl font-bold mb-1">{{ stock.Ticker }}</h2>
            <p class="text-slate-300 text-lg">{{ stock.Company }}</p>
          </div>
          <button @click="showDetails = false" class="text-slate-300 hover:text-white transition-colors">
            <X class="w-6 h-6" />
          </button>
        </div>
      </div>

      <div class="p-6 space-y-6 flex flex-col gap-4">
        <!-- Price Analysis -->
        <div
          class="bg-gradient-to-r from-green-50 to-blue-50 rounded-xl p-5 border border-gray-200 flex flex-col gap-2">
          <h3 class="text-lg font-semibold text-gray-800 mb-4 flex items-center">
            <Target class="w-5 h-5 mr-2 text-blue-600" />
            Price Targets
          </h3>
          <div class="grid grid-cols-2 gap-4">
            <div class="text-center">
              <div class="text-2xl font-bold text-green-600">
                ${{ stock.TargetFrom.toFixed(2) }}
              </div>
              <div class="text-sm text-gray-600">Minimum Target</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-blue-600">${{ stock.TargetTo.toFixed(2) }}</div>
              <div class="text-sm text-gray-600">Maximum Target</div>
            </div>
          </div>
          <div class="mt-4 text-center">
            <div class="text-lg font-semibold text-gray-700">
              Potential: {{ calculatePotential(stock) }}%
            </div>
          </div>
        </div>

        <!-- Rating Change -->
        <div class="bg-white rounded-xl p-5 border border-gray-200 shadow-sm">
          <h3 class="text-lg font-semibold text-gray-800 mb-4 flex items-center pb-2">
            Rating Update
          </h3>
          <div class="flex items-center justify-between space-x-8 gap-2">
            <div class="text-center m-4 w-32">
              <div class="bg-gray-100 rounded-lg px-4 py-2">
                <div class="text-sm text-gray-600">From</div>
                <div class="font-semibold text-gray-800">{{ stock.RatingFrom }}</div>
              </div>
            </div>
            <ArrowRight class="w-6 h-6 text-blue-500" />
            <div class="text-center m-4 w-32">
              <div class="bg-blue-100 rounded-lg px-4 py-2">
                <div class="text-sm text-blue-600">To</div>
                <div class="font-semibold text-blue-800">{{ stock.RatingTo }}</div>
              </div>
            </div>
          </div>
        </div>

        <!-- Action & Brokerage -->
        <div class="grid grid-cols-2 gap-4">
          <div class="bg-white rounded-xl p-4 border border-gray-200 shadow-sm">
            <h2 class="text-sm text-gray-600 mb-2">Action</h2>
            <p class="font-semibold text-gray-800">{{ stock.Action }}</p>
          </div>
          <div class="bg-white rounded-xl p-4 border border-gray-200 shadow-sm">
            <h2 class="text-sm text-gray-600 mb-2">Analyst</h2>
            <p class="font-semibold text-gray-800">{{ stock.Brokerage }}</p>
          </div>
        </div>

        <!-- Timestamp -->
        <div class="bg-gray-50 rounded-xl p-4 text-center">
          <div class="flex items-center justify-center text-gray-600">
            <Clock class="w-4 h-4 mr-2" />
            <span class="text-sm">Last updated {{ formatTime(stock.Time) }}</span>
          </div>
        </div>

        <!-- Back button -->
        <BaseButton @click="showDetails = false" class="w-full" variant="secondary">
          <ArrowLeft class="w-4 h-4 mr-2" />
          Close
        </BaseButton>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import BaseButton from '@/components/atoms/BaseButton.vue'
import type { Stock } from '@/types/stock'
import { ref } from 'vue'
const props = defineProps<{ stock: Stock }>()

import {
  ArrowDown,
  ArrowLeft,
  ArrowRight,
  ArrowUp,
  ChevronRight,
  Clock,
  Minus,
  Star,
  Target,
  TrendingUp,
  X,
} from 'lucide-vue-next'

// Reactive state
const showDetails = ref(false)

// Helper function to format time
const formatTime = (timeString: string) => {
  try {
    const date = new Date(timeString)
    return date.toLocaleDateString('es-ES', {
      day: '2-digit',
      month: '2-digit',
      year: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    })
  } catch {
    return timeString
  }
}

const formatTimeAgo = (timeString: string) => {
  try {
    const date = new Date(timeString)
    const now = new Date()
    const diffInHours = Math.floor((now.getTime() - date.getTime()) / (1000 * 60 * 60))

    if (diffInHours < 1) return 'less than 1h ago'
    if (diffInHours < 24) return `${diffInHours}h ago`
    const diffInDays = Math.floor(diffInHours / 24)
    return `${diffInDays}d ago`
  } catch {
    return 'recently'
  }
}

// Helper functions
const getActionColor = (action: string) => {
  switch (action.toLowerCase()) {
    case 'buy':
      return 'bg-green-100 text-green-800 border border-green-200'
    case 'sell':
      return 'bg-red-100 text-red-800 border border-red-200'
    case 'hold':
      return 'bg-yellow-100 text-yellow-800 border border-yellow-200'
    case 'top pick':
      return 'bg-green-100 text-green-800 border border-green-200'
    default:
      return 'bg-gray-100 text-gray-800 border border-gray-200'
  }
}

const getActionIcon = (action: string) => {
  switch (action.toLowerCase()) {
    case 'buy':
      return ArrowUp
    case 'sell':
      return ArrowDown
    case 'hold':
      return Minus
    case 'top pick':
      return Star
    default:
      return undefined
  }
}

const calculatePotential = (stock: Stock) => {
  const avg = (stock.TargetFrom + stock.TargetTo) / 2
  const current = stock.TargetFrom // Assuming current price is the lower target for demo
  return (((avg - current) / current) * 100).toFixed(1)
}
</script>
