package usecase

import (
	"JAVegaG/StockRecommendationAPI/core/domain"
	"JAVegaG/StockRecommendationAPI/core/service"
)

type RecommendTopStocksUseCase struct {
	stockService *service.StockService
}

func NewRecommendTopStocksUseCase(svc *service.StockService) *RecommendTopStocksUseCase {
	return &RecommendTopStocksUseCase{stockService: svc}
}

func (useCase *RecommendTopStocksUseCase) Execute() ([]*domain.Stock, error) {
	return useCase.stockService.GetRecommendations()
}
