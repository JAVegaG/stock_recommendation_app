package domain

import (
	"time"
)

type Stock struct {
	ID         string // uuid from google uuid module/package
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
