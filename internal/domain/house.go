package domain

import "context"

type House struct {
	ID          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      int     `json:"user_id" gorm:"not null"`
	HouseNumber string  `json:"house_number" gorm:"unique;not null"`
	Name        string  `json:"name" gorm:"not null"`
	Category    string  `json:"category" gorm:"not null"`
	Services    string  `json:"services" gorm:"not null"`
	Costs       float64 `json:"costs" gorm:"not null"`
}

type HouseRepository interface {
	Create(ctx context.Context, house *House) error
}
