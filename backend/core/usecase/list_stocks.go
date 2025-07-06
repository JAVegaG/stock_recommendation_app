package usecase

import (
	"JAVegaG/StockRecommendationAPI/core/domain"
	"JAVegaG/StockRecommendationAPI/core/service"
)

type ListStocksUseCase struct {
	stockService *service.StockService
}

func NewListStocksUseCase(svc *service.StockService) *ListStocksUseCase {
	return &ListStocksUseCase{stockService: svc}
}

func (useCase *ListStocksUseCase) Execute(limit int) ([]*domain.Stock, error) {
	recs, err := useCase.stockService.GetRecentStocks(limit)
	if err != nil {
		return nil, err
	}

	return recs, nil
}
