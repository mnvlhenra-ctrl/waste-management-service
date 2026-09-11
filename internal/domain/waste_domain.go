package domain

import (
	"context"
	"time"
)

type Waste struct {
	ID                int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name              string    `json:"name"`
	StockAvailability float64   `json:"stock_availability"`
	Costs             float64   `json:"costs"`
	Category          string    `json:"category"`
	IsSeparated       bool      `json:"is_separated" gorm:"default:false"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type WasteRepository interface {
	Create(ctx context.Context, waste *Waste) error
	// kutambahin GetAll buat parameter untuk filter ya mas
	GetAll(ctx context.Context, category string, isSeparated string) ([]Waste, error)
	GetByID(ctx context.Context, id int) (*Waste, error)
	Update(ctx context.Context, waste *Waste) error
	Delete(ctx context.Context, id int) error
}

type WasteUsecase interface {
	Create(ctx context.Context, waste *Waste) error
	// kutambahinGetAll buat parameter untuk filter ya mas
	GetAll(ctx context.Context, category string, isSeparated string) ([]Waste, error)
	GetByID(ctx context.Context, id int) (*Waste, error)
	Update(ctx context.Context, waste *Waste) error
	Delete(ctx context.Context, id int) error
}
