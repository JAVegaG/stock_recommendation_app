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

export async function getAll(limit?: number, offset: number = 0): Promise<IStockListResponse> {
  const url = `${endpoints.get}?offset=${offset}`

  const finalURL = limit ? `${url}&limit=${limit}` : url

  const res = await fetch(finalURL)
  const json = await res.json()

  return json.data as IStockListResponse
}
