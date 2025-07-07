package service

import (
	"JAVegaG/StockRecommendationAPI/core/domain"
	"JAVegaG/StockRecommendationAPI/core/repository"
)

type StockService struct {
	repo repository.StockRepository
}

func NewStockService(repo repository.StockRepository) *StockService {
	return &StockService{repo: repo}
}

func (s *StockService) SaveStock(stock *domain.Stock) error {
	return s.repo.Save(stock)
}

func (s *StockService) GetRecentStocks(limit int, offset int) ([]*domain.Stock, error) {
	return s.repo.FindRecent(limit, offset)
}

func (s *StockService) GetRecommendations(limit int, offset int) ([]*domain.Stock, error) {
	return s.repo.FindRecommendations(limit, offset)
}
