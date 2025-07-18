package handler

import (
	"JAVegaG/StockRecommendationAPI/core/domain"
	"JAVegaG/StockRecommendationAPI/core/usecase"
	"JAVegaG/StockRecommendationAPI/utils"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
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

	listQueryParams := getListQueryParams(&query)
	filterOptions, err := getFilterQueryParams(&query)

	if err != nil {
		utils.Logger.InfoContext(request.Context(), "Error getting filters from query params", slog.String("error", err.Error()))
		http.Error(response, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	stocks, err := handler.listStocksUC.Execute(listQueryParams.limit, listQueryParams.offset, filterOptions)
	if err != nil {
		utils.Logger.InfoContext(request.Context(), "Error getting stocks", slog.String("error", err.Error()))
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json") // Set Content-Type
	json.NewEncoder(response).Encode(stocks)
}

func (handler *StockHandler) GetRecommendations(response http.ResponseWriter, request *http.Request) {

	query := request.URL.Query()

	listQueryParams := getListQueryParams(&query)
	filterOptions, err := getFilterQueryParams(&query)
	if err != nil {
		utils.Logger.InfoContext(request.Context(), "Error getting filters from query params", slog.String("error", err.Error()))
		http.Error(response, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	recs, err := handler.recommendationUC.Execute(listQueryParams.limit, listQueryParams.offset, filterOptions)
	if err != nil {
		utils.Logger.InfoContext(request.Context(), "Error getting stock recommendations", slog.String("error", err.Error()))
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json") // Set Content-Type

	json.NewEncoder(response).Encode(recs)
}

type listQueryParams struct {
	limit  int
	offset int
}

func getListQueryParams(queryObject *url.Values) listQueryParams {
	limitStr := queryObject.Get("limit")
	limit := 10 // default

	l, err := strconv.Atoi(limitStr)

	if err == nil {
		limit = l
	}

	offsetStr := queryObject.Get("offset")
	offset := 0 // default

	o, err := strconv.Atoi(offsetStr)

	if err == nil {
		offset = o
	}

	return listQueryParams{
		limit,
		offset,
	}
}

func getFilterQueryParams(queryObject *url.Values) (*domain.StockFilterOptions, error) {
	var (
		Company     string
		TargetToMin float64
		TargetToMax float64
		RatingTo    string
	)

	Company = queryObject.Get("company")
	RatingTo = queryObject.Get("rating-to")
	targetToMinStr := queryObject.Get("target-to-min")
	targetToMaxStr := queryObject.Get("target-to-max")

	if targetToMinStr != "" {
		ttmin, err := strconv.ParseFloat(targetToMinStr, 64)

		if err == nil {
			TargetToMin = ttmin
		} else {
			return nil, err
		}
	}

	if targetToMaxStr != "" {
		ttmax, err := strconv.ParseFloat(targetToMaxStr, 64)

		if err == nil {
			TargetToMax = ttmax
		} else {
			return nil, err
		}
	}

	return &domain.StockFilterOptions{
		Company:     Company,
		TargetToMin: TargetToMin,
		TargetToMax: TargetToMax,
		RatingTo:    RatingTo,
	}, nil
}
