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

func (useCase *RecommendTopStocksUseCase) Execute(limit int, offset int, filterOptions *domain.StockFilterOptions) (*domain.StockListResponse, error) {
	return useCase.stockService.GetRecommendations(limit, offset, filterOptions)
}
