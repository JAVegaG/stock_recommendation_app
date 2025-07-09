export interface IEndpoints {
  get: string
  getRecommendations: string
}

export interface IFilterOptions {
  Company?: string
  TargetToMin?: number
  TargetToMax?: number
  RatingTo?: string
}

export const ratings = [
  'Buy',
  'Sector Perform',
  'Overweight',
  'Neutral',
  'Equal Weight',
  'Underweight',
  'Outperform',
  'Sell',
  'Hold',
  'Market Perform',
  'Speculative Buy',
  'Market Outperform',
  'Underperform',
  'Sector Outperform',
  'Strong-Buy',
  'Positive',
  'Sector Underperform',
  'Mkt Outperform',
  'Moderate Buy',
  'In-Line',
  'Strong Sell',
  'Sector Weight',
  'Top Pick',
  'Outperformer',
  'Reduce',
] as const
