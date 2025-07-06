package repository

import "JAVegaG/StockRecommendationAPI/core/domain"

type StockRepository interface {
	Save(stock *domain.Stock) error
	FindRecent(limit int) ([]*domain.Stock, error)
	FindRecommendations() ([]*domain.Stock, error)
}
