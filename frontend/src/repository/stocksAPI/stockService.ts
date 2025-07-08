import type { IFilterOptions } from './interfaces/settings'
import type { IStockListResponse } from './interfaces/stocks'
import { endpoints } from './settings'

export async function getRecommendations(
  limit?: number,
  offset: number = 0,
): Promise<IStockListResponse> {
  const url = `${endpoints.getRecommendations}?offset=${offset}`

  const finalURL = limit ? `${url}&limit=${limit}` : url

  const res = await fetch(finalURL)
  const json = await res.json()

  return json.data as IStockListResponse
}

export async function getAll(limit?: number, offset: number = 0, filterOptions?: IFilterOptions): Promise<IStockListResponse> {
  let url = `${endpoints.get}?offset=${offset}`

  if (limit) {
    url += `&limit=${limit}`
  }

  if (filterOptions?.Company) {
    url += `&company=${filterOptions?.Company}`
  }

  if (filterOptions?.RatingTo) {
    url += `&rating-to=${filterOptions?.RatingTo}`
  }

  if (filterOptions?.TargetToMin) {
    url += `&target-to-min=${filterOptions?.TargetToMin}`
  }

  if (filterOptions?.TargetToMax) {
    url += `&target-to-max=${filterOptions?.TargetToMax}`
  }

  const res = await fetch(url)
  const json = await res.json()

  return json.data as IStockListResponse
}
