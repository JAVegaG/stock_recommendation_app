import type { Stock } from '@/types/stock'

export interface IStockListResponse {
  Items: Stock[]
  Total: number
}
