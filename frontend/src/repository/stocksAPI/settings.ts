import type { IEndpoints } from "./interfaces/settings"

export const baseURL: string = import.meta.env.VITE_API_URL

export const endpoints: IEndpoints = {
    get: `${baseURL}/api/stocks/`,
    getRecommendations: `${baseURL}/api/stocks/recommendations`
}