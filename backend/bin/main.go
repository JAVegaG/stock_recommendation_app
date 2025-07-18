package main

import (
	"JAVegaG/StockRecommendationAPI/core/service"
	"JAVegaG/StockRecommendationAPI/core/usecase"
	"JAVegaG/StockRecommendationAPI/infra/api"
	"JAVegaG/StockRecommendationAPI/infra/db/cockroach"
	"JAVegaG/StockRecommendationAPI/infra/handler"
	responseMiddleware "JAVegaG/StockRecommendationAPI/infra/middleware"
	"JAVegaG/StockRecommendationAPI/utils"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/traceid"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Printf("Error loading .env file %v\nNow using env vars\n", err)
	}

	db, err := cockroach.NewConnection()
	if err != nil {
		log.Fatalf("Error connecting to CockroachDB: %v", err)
	}

	stockRepo := cockroach.NewStockRepository(db)

	stockService := service.NewStockService(stockRepo)

	apiClient := api.NewStockAPIClient()

	listUC := usecase.NewListStocksUseCase(stockService)
	recommendUC := usecase.NewRecommendTopStocksUseCase(stockService)
	storeUC := usecase.NewStoreStocksFromAPIUseCase(apiClient, stockService)

	stockHandler := handler.NewStockHandler(recommendUC, listUC)

	router := chi.NewRouter()
	router.Use(traceid.Middleware)
	router.Use(utils.HttpRequestLogger)
	router.Use(middleware.Recoverer)
	router.Use(responseMiddleware.ResponseWrapper)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{os.Getenv("CORS_ORIGIN")},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	stockHandler.RegisterRoutes(router)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	if strings.EqualFold(os.Getenv("SEED_API"), "true") {
		err := storeUC.Execute()
		if err != nil {
			log.Fatalf("Error loading data from external API: %v", err)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server Listening on Port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
