package usecase

import (
	"JAVegaG/StockRecommendationAPI/core/domain"
	"JAVegaG/StockRecommendationAPI/core/service"
	"JAVegaG/StockRecommendationAPI/infra/api"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type StoreStocksFromAPIUseCase struct {
	apiClient    api.StockAPIClient
	stockService *service.StockService
}

func NewStoreStocksFromAPIUseCase(apiClient api.StockAPIClient, svc *service.StockService) *StoreStocksFromAPIUseCase {
	return &StoreStocksFromAPIUseCase{
		apiClient:    apiClient,
		stockService: svc,
	}
}

func (useCase *StoreStocksFromAPIUseCase) Execute() error {
	nextPage := ""

	for {
		data, err := useCase.apiClient.FetchPage(nextPage)

		if err != nil {
			return err
		}

		for _, item := range data.Items {
			parsedTime, _ := time.Parse(time.RFC3339Nano, item.Time)
			targetFrom := parseDollar(item.TargetFrom)
			targetTo := parseDollar(item.TargetTo)

			stock := &domain.Stock{
				ID:         uuid.NewString(),
				Ticker:     item.Ticker,
				Company:    item.Company,
				TargetFrom: targetFrom,
				TargetTo:   targetTo,
				Action:     item.Action,
				Brokerage:  item.Brokerage,
				RatingFrom: item.RatingFrom,
				RatingTo:   item.RatingTo,
				Time:       parsedTime,
			}

			useCase.stockService.SaveStock(stock)
		}

		if data.NextPage == "" {
			break
		}

		nextPage = data.NextPage
	}

	return nil
}

func parseDollar(s string) float64 {
	var val float64
	fmt.Sscanf(strings.ReplaceAll(s, "$", ""), "%f", &val)
	return val
}
