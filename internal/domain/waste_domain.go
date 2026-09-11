package domain

import (
	"context"
	"time"
)

type Waste struct {
	ID                int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name              string    `json:"name"`
	StockAvailability float64   `json:"stock_availability"`
	SeparatedCosts    float64   `json:"separated_costs"`
	NonSeparatedCosts float64   `json:"non_separated_costs"`
	Category          string    `json:"category"`
	IsSeparated       bool      `json:"is_separated" gorm:"default:false"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type WasteRepository interface {
	Create(ctx context.Context, waste *Waste) error
	GetAll(ctx context.Context, category string, isSeparated string) ([]Waste, error)
	GetByID(ctx context.Context, id int) (*Waste, error)
	Update(ctx context.Context, waste *Waste) error
	Delete(ctx context.Context, id int) error
}

type WasteUsecase interface {
	Create(ctx context.Context, waste *Waste) error
	GetAll(ctx context.Context, category string, isSeparated string) ([]Waste, error)
	GetByID(ctx context.Context, id int) (*Waste, error)
	Update(ctx context.Context, waste *Waste) error
	Delete(ctx context.Context, id int) error
}
