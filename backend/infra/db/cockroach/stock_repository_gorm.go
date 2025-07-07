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

func (r *gormStockRepository) FindAll(limit int, offset int, filterOptions *domain.StockFilterOptions) (*domain.StockListResponse, error) {
	var models []StockModel
	var total int64

	countResult := r.db.Model(&StockModel{})
	queryResult := r.db.Order("time desc")

	if filterOptions.Company != "" {
		countResult = countResult.Where("company LIKE ?", filterOptions.Company+"%")
		queryResult = queryResult.Where("company LIKE ?", filterOptions.Company+"%")
	}

	if filterOptions.RatingTo != "" {
		countResult = countResult.Where("rating_to = ?", filterOptions.RatingTo)
		queryResult = queryResult.Where("rating_to = ?", filterOptions.RatingTo)
	}

	if filterOptions.TargetToMin != 0 {
		countResult = countResult.Where("target_to >= ?", filterOptions.Company)
		queryResult = queryResult.Where("target_to >= ?", filterOptions.TargetToMin)
	}

	if filterOptions.TargetToMax != 0 {
		countResult = countResult.Where("target_to <= ?", filterOptions.TargetToMax)
		queryResult = queryResult.Where("target_to <= ?", filterOptions.TargetToMax)
	}

	countResult = countResult.Count(&total)

	err := countResult.Error

	if err != nil {
		return nil, err
	}

	queryResult = queryResult.Offset(offset).Limit(limit).Find(&models)

	err = queryResult.Error

	if err != nil {
		return nil, err
	}

	result := domain.StockListResponse{
		Items: toDomainList(models),
		Total: total,
	}

	return (&result), nil
}

func (r *gormStockRepository) FindRecommendations(limit int, offset int, filterOptions *domain.StockFilterOptions) (*domain.StockListResponse, error) {
	var models []StockModel
	var total int64

	countQuery := `
    SELECT COUNT(*)
    FROM stock_models
    WHERE (rating_from != rating_to
      AND (
        CASE rating_to
          WHEN 'Strong Sell' THEN 0
          WHEN 'Sell' THEN 1
          WHEN 'Underperform' THEN 2
          WHEN 'Reduce' THEN 3
          WHEN 'Hold' THEN 4
          WHEN 'Neutral' THEN 5
          WHEN 'Equal Weight' THEN 6
          WHEN 'Market Perform' THEN 7
          WHEN 'In-Line' THEN 8
          WHEN 'Sector Weight' THEN 9
          WHEN 'Sector Perform' THEN 10
          WHEN 'Outperform' THEN 11
          WHEN 'Market Outperform' THEN 12
          WHEN 'Overweight' THEN 13
          WHEN 'Sector Outperform' THEN 14
          WHEN 'Outperformer' THEN 15
          WHEN 'Positive' THEN 16
          WHEN 'Mkt Outperform' THEN 17
          WHEN 'Speculative Buy' THEN 18
          WHEN 'Moderate Buy' THEN 19
          WHEN 'Buy' THEN 20
          WHEN 'Strong-Buy' THEN 21
          WHEN 'Top Pick' THEN 22
          ELSE -1
        END
      )
      >
      (
        CASE rating_from
          WHEN 'Strong Sell' THEN 0
          WHEN 'Sell' THEN 1
          WHEN 'Underperform' THEN 2
          WHEN 'Reduce' THEN 3
          WHEN 'Hold' THEN 4
          WHEN 'Neutral' THEN 5
          WHEN 'Equal Weight' THEN 6
          WHEN 'Market Perform' THEN 7
          WHEN 'In-Line' THEN 8
          WHEN 'Sector Weight' THEN 9
          WHEN 'Sector Perform' THEN 10
          WHEN 'Outperform' THEN 11
          WHEN 'Market Outperform' THEN 12
          WHEN 'Overweight' THEN 13
          WHEN 'Sector Outperform' THEN 14
          WHEN 'Outperformer' THEN 15
          WHEN 'Positive' THEN 16
          WHEN 'Mkt Outperform' THEN 17
          WHEN 'Speculative Buy' THEN 18
          WHEN 'Moderate Buy' THEN 19
          WHEN 'Buy' THEN 20
          WHEN 'Strong-Buy' THEN 21
          WHEN 'Top Pick' THEN 22
          ELSE -1
        END
	)) OR rating_to = 'Top Pick'
    `

	countResult := r.db.Raw(countQuery).Scan(&total)

	err := countResult.Error

	if err != nil {
		return nil, err
	}

	query := `
    SELECT *
    FROM stock_models
    WHERE (rating_from != rating_to
      AND (
        CASE rating_to
          WHEN 'Strong Sell' THEN 0
          WHEN 'Sell' THEN 1
          WHEN 'Underperform' THEN 2
          WHEN 'Reduce' THEN 3
          WHEN 'Hold' THEN 4
          WHEN 'Neutral' THEN 5
          WHEN 'Equal Weight' THEN 6
          WHEN 'Market Perform' THEN 7
          WHEN 'In-Line' THEN 8
          WHEN 'Sector Weight' THEN 9
          WHEN 'Sector Perform' THEN 10
          WHEN 'Outperform' THEN 11
          WHEN 'Market Outperform' THEN 12
          WHEN 'Overweight' THEN 13
          WHEN 'Sector Outperform' THEN 14
          WHEN 'Outperformer' THEN 15
          WHEN 'Positive' THEN 16
          WHEN 'Mkt Outperform' THEN 17
          WHEN 'Speculative Buy' THEN 18
          WHEN 'Moderate Buy' THEN 19
          WHEN 'Buy' THEN 20
          WHEN 'Strong-Buy' THEN 21
          WHEN 'Top Pick' THEN 22
          ELSE -1
        END
      )
      >
      (
        CASE rating_from
          WHEN 'Strong Sell' THEN 0
          WHEN 'Sell' THEN 1
          WHEN 'Underperform' THEN 2
          WHEN 'Reduce' THEN 3
          WHEN 'Hold' THEN 4
          WHEN 'Neutral' THEN 5
          WHEN 'Equal Weight' THEN 6
          WHEN 'Market Perform' THEN 7
          WHEN 'In-Line' THEN 8
          WHEN 'Sector Weight' THEN 9
          WHEN 'Sector Perform' THEN 10
          WHEN 'Outperform' THEN 11
          WHEN 'Market Outperform' THEN 12
          WHEN 'Overweight' THEN 13
          WHEN 'Sector Outperform' THEN 14
          WHEN 'Outperformer' THEN 15
          WHEN 'Positive' THEN 16
          WHEN 'Mkt Outperform' THEN 17
          WHEN 'Speculative Buy' THEN 18
          WHEN 'Moderate Buy' THEN 19
          WHEN 'Buy' THEN 20
          WHEN 'Strong-Buy' THEN 21
          WHEN 'Top Pick' THEN 22
          ELSE -1
        END
	)) OR rating_to = 'Top Pick'
    ORDER BY time DESC
    LIMIT ? OFFSET ?;
    `

	queryResult := r.db.Raw(query, limit, offset).Scan(&models)

	err = queryResult.Error

	if err != nil {
		return nil, err
	}

	result := domain.StockListResponse{
		Items: toDomainList(models),
		Total: total,
	}

	return (&result), nil
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
