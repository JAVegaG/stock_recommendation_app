import type { Stock } from '@/types/stock'
import { endpoints } from './settings'

export async function getRecommendations(limit?: number, offset?: number): Promise<Stock[]> {


    const res = await fetch(`${endpoints.getRecommendations}`)
    console.log({res})
    const json = await res.json()

    console.log(json)

    return json.data as Stock[]
}
