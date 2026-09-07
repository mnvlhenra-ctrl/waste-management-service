package domain

import "time"

type Transaction struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Amount          float64   `json:"amount"`
	UserID          int       `json:"user_id"`
	InvoiceID       *int      `json:"invoice_id,omitempty"`
	TransactionType string    `json:"transaction_type"`
	TransactionDate time.Time `json:"transaction_date"`
	Status          string    `json:"status"`
}
