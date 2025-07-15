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
	// 1. Cargar .env
	err := godotenv.Load()

	if err != nil {
		log.Println("No se cargó archivo .env, usando variables de entorno")
	}

	// 2. Conectar a la base de datos
	db, err := cockroach.NewConnection()
	if err != nil {
		log.Fatalf("Error al conectar a CockroachDB: %v", err)
	}

	// 3. Instanciar repositorio GORM
	stockRepo := cockroach.NewStockRepository(db)

	// 4. Crear servicio de dominio
	stockService := service.NewStockService(stockRepo)

	// 4. Crear clinete de api externa
	apiClient := api.NewStockAPIClient()

	// 5. Crear casos de uso
	listUC := usecase.NewListStocksUseCase(stockService)
	recommendUC := usecase.NewRecommendTopStocksUseCase(stockService)
	storeUC := usecase.NewStoreStocksFromAPIUseCase(apiClient, stockService)

	// 6. Crear handlers y router
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

	// 7. Registrar rutas
	stockHandler.RegisterRoutes(router)

	// Ejecutar solo si quieres precargar
	if strings.EqualFold(os.Getenv("SEED_API"), "true") {
		err := storeUC.Execute()
		if err != nil {
			log.Fatalf("Error cargando datos desde el API externo: %v", err)
		}
	}

	// 8. Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Servidor escuchando en http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
