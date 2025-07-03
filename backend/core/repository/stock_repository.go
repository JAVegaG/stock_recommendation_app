package repository

import "JAVegaG/StockRecommendationAPI/core/domain"

type StockRepository interface {
	Save(stock *domain.Stock) error
	FindAll(limit int, offset int, filterOptions *domain.StockFilterOptions) (*domain.StockListResponse, error)
	FindRecommendations(limit int, offset int, filterOptions *domain.StockFilterOptions) (*domain.StockListResponse, error)
}
