<template>
  <div class="max-w-sm mx-auto">
    <div
      class="bg-gradient-to-br from-white to-gray-50 rounded-2xl shadow-lg hover:shadow-xl transition-all duration-300 border border-gray-200 overflow-hidden h-[27rem] flex flex-col gap-2"
    >
      <!-- Header with Ticker -->
      <div class="bg-gray-700 p-4 text-white">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="bg-gray-600 h-9 w-[9rem]"></h3>
            <p class="opacity-90 bg-gray-500 h-5 w-[9rem]"></p>
          </div>
          <div class="bg-white/20 rounded-full p-2"></div>
        </div>
      </div>

      <!-- Price Range -->
      <div class="p-6 space-y-4 flex flex-col gap-4 h-full justify-between">
        <div class="grid grid-cols-2 gap-4">
          <div
            class="bg-gray-50 border border-gray-200 rounded-xl p-4 text-center flex flex-col items-center"
          >
            <div class="text-gray-600 text-xs font-semibold uppercase tracking-wide mb-1">
              Target Min
            </div>
            <div class="bg-gray-600 min-h-5 w-[7.5rem]"></div>
          </div>
          <div
            class="bg-gray-50 border border-gray-200 rounded-xl p-4 text-center flex flex-col items-center"
          >
            <div class="text-gray-600 text-xs font-semibold uppercase tracking-wide mb-1">
              Target Max
            </div>
            <div class="bg-gray-600 min-h-5 w-[7.5rem]"></div>
          </div>
        </div>

        <!-- Action Badge -->
        <div class="flex justify-center">
          <span
            class="inline-flex items-center justify-center gap-2 px-4 py-2 rounded-full bg-gray-600 w-[11.5rem]"
          >
            <div class="w-4 h-4 mr-2 bg-gray-400"></div>
            <div class="bg-gray-300 min-h-5 w-[6rem]"></div>
          </span>
        </div>

        <!-- Update Info -->
        <div
          class="flex gap-8 items-center justify-between text-sm text-gray-600 bg-gray-50 rounded-lg p-3"
        >
          <div class="flex items-center w-32 gap-2">
            <div class="w-4 h-4 mr-2 bg-gray-400"></div>
            <div class="flex flex-col">
              Updated
              <div class="bg-gray-300 min-h-5 w-[6rem]"></div>
            </div>
          </div>
          <div class="flex flex-col">
            Brokerage
            <div class="min-h-5 w-32 bg-gray-300"></div>
          </div>
        </div>
      </div>

      <!-- Ver más button -->
      <div class="px-6 pb-6" style="margin-top: auto">
        <BaseButton
          class="w-full group bg-gray-600"
          :disabled="true"
          :loading="true"
          :variant="'secondary'"
        >
          <span class="flex items-center justify-center">
            View details
            <ChevronRight class="w-4 h-4 ml-2" />
          </span>
        </BaseButton>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import BaseButton from '@/components/atoms/BaseButton.vue'
import type { Stock } from '@/types/stock'

import { ChevronRight } from 'lucide-vue-next'
</script>
