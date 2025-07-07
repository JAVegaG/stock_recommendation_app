package repository

import "JAVegaG/StockRecommendationAPI/core/domain"

type StockRepository interface {
	Save(stock *domain.Stock) error
	FindRecent(limit int, offset int) ([]*domain.Stock, error)
	FindRecommendations(limit int, offset int) ([]*domain.Stock, error)
}
