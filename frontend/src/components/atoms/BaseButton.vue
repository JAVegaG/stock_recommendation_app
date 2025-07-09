<template>
  <button
    :class="buttonClasses"
    :disabled="disabled || loading"
    @click="handleClick"
    v-bind="$attrs"
  >
    <span v-if="loading" class="flex items-center justify-center gap-2">
      <svg
        class="animate-spin -ml-1 mr-3 h-4 w-4"
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
      >
        <circle
          class="opacity-25"
          cx="12"
          cy="12"
          r="10"
          stroke="currentColor"
          stroke-width="4"
        ></circle>
        <path
          class="opacity-75"
          fill="currentColor"
          d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
        ></path>
      </svg>
      {{ loadingText || 'Loading...' }}
    </span>
    <span v-else class="flex items-center justify-center">
      <component v-if="iconLeft" :is="iconLeft" :class="iconClasses" />
      <slot />
      <component v-if="iconRight" :is="iconRight" :class="iconClasses" />
    </span>
  </button>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { Props } from './BaseButton.interface'

const props = withDefaults(defineProps<Props>(), {
  variant: 'primary',
  size: 'md',
  disabled: false,
  loading: false,
  fullWidth: false,
  rounded: false,
})

const emit = defineEmits<{
  click: [event: MouseEvent]
}>()

const baseClasses =
  'inline-flex items-center justify-center font-semibold transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 transform active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none'

const variantClasses = computed(() => {
  const variants = {
    primary:
      'bg-gradient-to-r from-blue-600 to-blue-700 hover:from-blue-700 hover:to-blue-800 text-white shadow-lg hover:shadow-xl focus:ring-blue-500',
    secondary:
      'bg-white hover:bg-gray-50 text-gray-700 border-2 border-gray-300 hover:border-gray-400 shadow-sm hover:shadow-md focus:ring-gray-500',
    success:
      'bg-gradient-to-r from-green-600 to-green-700 hover:from-green-700 hover:to-green-800 text-white shadow-lg hover:shadow-xl focus:ring-green-500',
    danger:
      'bg-gradient-to-r from-red-600 to-red-700 hover:from-red-700 hover:to-red-800 text-white shadow-lg hover:shadow-xl focus:ring-red-500',
    warning:
      'bg-gradient-to-r from-yellow-500 to-yellow-600 hover:from-yellow-600 hover:to-yellow-700 text-white shadow-lg hover:shadow-xl focus:ring-yellow-500',
    ghost: 'text-gray-600 hover:text-gray-800 hover:bg-gray-100 focus:ring-gray-500',
    outline:
      'border-2 border-blue-600 text-blue-600 hover:bg-blue-600 hover:text-white focus:ring-blue-500',
  }
  return variants[props.variant]
})

const sizeClasses = computed(() => {
  const sizes = {
    xs: 'px-2.5 py-1.5 text-xs',
    sm: 'px-3 py-2 text-sm',
    md: 'px-4 py-2.5 text-sm',
    lg: 'px-6 py-3 text-base',
    xl: 'px-8 py-4 text-lg',
  }
  return sizes[props.size]
})

const roundedClasses = computed(() => {
  if (props.rounded) return 'rounded-full'

  const roundedSizes = {
    xs: 'rounded-md',
    sm: 'rounded-md',
    md: 'rounded-lg',
    lg: 'rounded-xl',
    xl: 'rounded-xl',
  }
  return roundedSizes[props.size]
})

const widthClasses = computed(() => {
  return props.fullWidth ? 'w-full' : ''
})

const iconClasses = computed(() => {
  const iconSizes = {
    xs: 'w-3 h-3',
    sm: 'w-4 h-4',
    md: 'w-4 h-4',
    lg: 'w-5 h-5',
    xl: 'w-6 h-6',
  }

  const spacing = props.iconLeft ? 'mr-2' : props.iconRight ? 'ml-2' : ''
  return `${iconSizes[props.size]} ${spacing}`
})

const buttonClasses = computed(() => {
  return [
    baseClasses,
    variantClasses.value,
    sizeClasses.value,
    roundedClasses.value,
    widthClasses.value,
  ].join(' ')
})

const handleClick = (event: MouseEvent) => {
  if (!props.disabled && !props.loading) {
    emit('click', event)
  }
}
</script>
