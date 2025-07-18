import type { IFilterOptions } from './interfaces/settings'
import type { IStockListResponse } from './interfaces/stocks'
import { endpoints } from './settings'

export async function getRecommendations(
  limit?: number,
  offset: number = 0,
  filterOptions?: IFilterOptions,
): Promise<IStockListResponse> {
  let url = `${endpoints().getRecommendations}?offset=${offset}`

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

  try {

    const res = await fetch(url)

    if (res.status.toString().startsWith('2')) {
      const json = await res.json()

      return json.data as IStockListResponse
    }

    return {
      Items: [],
      Total: 0
    }

  } catch (e) {
    console.error({ error: e })
    return {
      Items: [],
      Total: 0
    }
  }
}

export async function getAll(
  limit?: number,
  offset: number = 0,
  filterOptions?: IFilterOptions,
): Promise<IStockListResponse> {
  let url = `${endpoints().get}?offset=${offset}`

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


  try {

    const res = await fetch(url)
    if (res.status.toString().startsWith('2')) {
      const json = await res.json()

      return json.data as IStockListResponse
    }

    return {
      Items: [],
      Total: 0
    }
  } catch (e) {

    console.error({ error: e })

    return {
      Items: [],
      Total: 0
    }
  }

}
