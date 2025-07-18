<template>
  <div class="max-w-sm mx-auto">
    <div
      class="bg-gradient-to-br from-white to-gray-50 rounded-2xl shadow-lg hover:shadow-xl transition-all duration-300 border border-gray-200 overflow-hidden h-[27rem] flex flex-col gap-2">
      <!-- Header with Ticker -->
      <div class="bg-gradient-to-r from-blue-600 to-blue-700 p-4 opacity-50 text-white">
        <div class="flex items-center justify-between">
          <div>
            <div class="bg-blue-400 h-8 w-[9rem] animate-pulse rounded-lg"></div>
            <div class="opacity-90 bg-blue-100 h-5 w-[9rem] animate-pulse rounded-lg"></div>
          </div>
          <div class="bg-white/20 rounded-full p-2 animate-pulse"></div>
        </div>
      </div>

      <!-- Price Range -->
      <div class="p-6 space-y-4 flex flex-col gap-4 h-full justify-between">
        <div class="grid grid-cols-2 gap-4">
          <div class="bg-green-50 border border-gray-200 rounded-xl p-4 text-center flex flex-col items-center">
            <div class="text-green-600 text-xs font-semibold uppercase tracking-wide mb-1">
              Target Min
            </div>
            <div class="bg-green-800 opacity-50 min-h-5 w-[6rem] animate-pulse rounded-lg"></div>
          </div>
          <div class="bg-blue-50 border border-gray-200 rounded-xl p-4 text-center flex flex-col items-center">
            <div class="text-blue-600 text-xs font-semibold uppercase tracking-wide mb-1">
              Target Max
            </div>
            <div class="bg-blue-800 opacity-50 min-h-5 w-[6rem] animate-pulse rounded-lg"></div>
          </div>
        </div>

        <!-- Action Badge -->
        <div class="flex justify-center">
          <span class="inline-flex items-center justify-center gap-2 px-4 py-2 rounded-full w-[11.5rem]">
            <div class="w-4 h-4 mr-2 bg-gray-400 animate-pulse rounded-lg"></div>
            <div class="bg-gray-300 min-h-5 w-[6rem] animate-pulse rounded-lg"></div>
          </span>
        </div>

        <!-- Update Info -->
        <div class="grid grid-cols-2 gap-4 items-start justify-between text-sm text-gray-600 bg-gray-50 rounded-lg p-3">
          <div class="flex gap-1">

            <div class="w-4 h-4 mr-2 bg-gray-400 animate-pulse rounded-lg"></div>
            <div class="flex flex-col">
              Updated
              <div class="bg-gray-300 min-h-5 w-[6rem] animate-pulse rounded-lg"></div>
            </div>

          </div>
          <div class="flex flex-col">
            Brokerage
            <div class="bg-gray-300 min-h-5 w-[7rem] animate-pulse rounded-lg"></div>
          </div>
        </div>
      </div>

      <!-- Ver más button -->
      <div class="px-6 pb-6" style="margin-top: auto">
        <BaseButton class="w-full group" :loading="true" :variant="'primary'">
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
import BaseButton from '@/components/atoms/BaseButton.vue';

import { ChevronRight } from 'lucide-vue-next';
</script>
