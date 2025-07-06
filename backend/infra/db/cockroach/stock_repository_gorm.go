package cockroach

import (
	"JAVegaG/StockRecommendationAPI/core/domain"
	"time"

	"gorm.io/gorm"
)

// Estructura para mapear con la tabla real
type StockModel struct {
	ID         string `gorm:"primaryKey"`
	Ticker     string
	Company    string
	TargetFrom float64
	TargetTo   float64
	Action     string
	Brokerage  string
	RatingFrom string
	RatingTo   string
	Time       time.Time
}

type gormStockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) *gormStockRepository {
	db.AutoMigrate(&StockModel{})
	return &gormStockRepository{db: db}
}

func (r *gormStockRepository) Save(stock *domain.Stock) error {
	model := toModel(stock)
	return r.db.Create(model).Error
}

func (r *gormStockRepository) FindRecent(limit int) ([]*domain.Stock, error) {
	var models []StockModel

	queryResult := r.db.Order("time desc").Limit(limit).Find(&models)

	err := queryResult.Error

	if err != nil {
		return nil, err
	}

	return toDomainList(models), nil
}

func (r *gormStockRepository) FindRecommendations() ([]*domain.Stock, error) {
	var models []StockModel

	queryResult := r.db.Where("rating_from != rating_to").Order("time desc").Limit(10).Find(&models)

	err := queryResult.Error

	if err != nil {
		return nil, err
	}

	return toDomainList(models), nil
}

// Mappers - DTOs
func toModel(stock *domain.Stock) *StockModel {
	return &StockModel{
		ID:         stock.ID,
		Ticker:     stock.Ticker,
		Company:    stock.Company,
		TargetFrom: stock.TargetFrom,
		TargetTo:   stock.TargetTo,
		Action:     stock.Action,
		Brokerage:  stock.Brokerage,
		RatingFrom: stock.RatingFrom,
		RatingTo:   stock.RatingTo,
		Time:       stock.Time,
	}
}

func toDomainList(models []StockModel) []*domain.Stock {
	stocks := make([]*domain.Stock, 0, len(models))

	for _, model := range models {
		stocks = append(stocks, &domain.Stock{
			ID:         model.ID,
			Ticker:     model.Ticker,
			Company:    model.Company,
			TargetFrom: model.TargetFrom,
			TargetTo:   model.TargetTo,
			Action:     model.Action,
			Brokerage:  model.Brokerage,
			RatingFrom: model.RatingFrom,
			RatingTo:   model.RatingTo,
			Time:       model.Time,
		})
	}
	return stocks
}
