import { getEnv } from '@/utils/env'
import type { IEndpoints } from './interfaces/settings'

export const baseURL: () => string = () => getEnv('VITE_API_URL') as string || ''

export const endpoints: () => IEndpoints = () => ({
  get: `${baseURL()}/api/stocks`,
  getRecommendations: `${baseURL()}/api/stocks/recommendations`,
})
