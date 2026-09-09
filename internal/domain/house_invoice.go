package domain

import (
	"context"
	"time"
)

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

type Invoice struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	HouseID      int       `json:"house_id" gorm:"not null"`
	BillingMonth string    `json:"billing_month" gorm:"not null"`
	Amount       float64   `json:"amount" gorm:"not null"`
	Status       string    `json:"status" gorm:"not null;default:'UNPAID'"`
	DueDate      time.Time `json:"due_date" gorm:"not null"`
}

type InvoiceRepository interface {
	Create(ctx context.Context, invoice *Invoice) error
	GetByID(ctx context.Context, id int) (*Invoice, error)
	GetByUserID(ctx context.Context, userID int) ([]Invoice, error)
	UpdateStatus(ctx context.Context, id int, status string) error
}
