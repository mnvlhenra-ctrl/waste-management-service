package domain

import (
	"context"
	"time"
)

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
	UpdateStatus(ctx context.Context, id int, status string) error
}
