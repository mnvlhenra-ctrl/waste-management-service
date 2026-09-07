package domain

import "time"

type Transaction struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	InvoiceID       *int      `json:"invoice_id"`
	Name            string    `json:"name"`
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	Status          string    `json:"status"`
	TransactionDate time.Time `json:"transaction_date"`
}
