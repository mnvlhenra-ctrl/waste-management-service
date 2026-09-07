package domain

import (
	"context"
	"time"
)

type Waste struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	StockAvailability float64   `json:"stock_availability"`
	Costs             float64   `json:"costs"`
	Category          string    `json:"category"`
	CreatedAt         time.Time `json:"created_at"`
}

type WasteRepository interface {
	Create(
		ctx context.Context,
		waste *Waste,
	) error

	GetAll(
		ctx context.Context,
	) ([]Waste, error)

	GetByID(
		ctx context.Context,
		id int,
	) (*Waste, error)

	Update(
		ctx context.Context,
		waste *Waste,
	) error

	Delete(
		ctx context.Context,
		id int,
	) error
}

type WasteUsecase interface {
	Create(
		ctx context.Context,
		waste *Waste,
	) error

	GetAll(
		ctx context.Context,
	) ([]Waste, error)

	GetByID(
		ctx context.Context,
		id int,
	) (*Waste, error)

	Update(
		ctx context.Context,
		waste *Waste,
	) error

	Delete(
		ctx context.Context,
		id int,
	) error
}
