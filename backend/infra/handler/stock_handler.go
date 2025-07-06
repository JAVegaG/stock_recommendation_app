package handler

import (
	"JAVegaG/StockRecommendationAPI/core/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type StockHandler struct {
	recommendationUC *usecase.RecommendTopStocksUseCase
	listStocksUC     *usecase.ListStocksUseCase
}

func NewStockHandler(recommendUC *usecase.RecommendTopStocksUseCase, listUC *usecase.ListStocksUseCase) *StockHandler {
	return &StockHandler{
		recommendationUC: recommendUC,
		listStocksUC:     listUC,
	}
}

func (handler *StockHandler) RegisterRoutes(routes chi.Router) {
	routes.Get("/api/stocks", handler.GetStocks)
	routes.Get("/api/stocks/recommendations", handler.GetRecommendations)
}

func (handler *StockHandler) GetStocks(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	limitStr := query.Get("limit")
	limit := 20 // default

	l, err := strconv.Atoi(limitStr)

	if err == nil {
		limit = l
	}

	stocks, err := handler.listStocksUC.Execute(limit)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json") // Set Content-Type
	json.NewEncoder(response).Encode(stocks)
}

func (handler *StockHandler) GetRecommendations(response http.ResponseWriter, request *http.Request) {
	recs, err := handler.recommendationUC.Execute()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json") // Set Content-Type

	json.NewEncoder(response).Encode(recs)
}
