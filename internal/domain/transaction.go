package domain

import (
	"context"
)

type Transaction struct {
	ID              int     `json:"id"`
	UserID          int     `json:"user_id"`
	InvoiceID       *int    `json:"invoice_id"`
	Name            string  `json:"name"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	Status          string  `json:"status"`
	TransactionDate string  `json:"transaction_date"`
}

type TransactionRepository interface {
	Create(
		ctx context.Context,
		transaction *Transaction,
	) error

	GetAllByUserID(
		ctx context.Context,
		userID int,
	) ([]Transaction, error)

	GetByID(
		ctx context.Context,
		id int,
	) (*Transaction, error)

	UpdateStatus(
		ctx context.Context,
		id int,
		status string,
	) error

	Delete(
		ctx context.Context,
		id int,
	) error

	UpdateUserBalance(
		ctx context.Context,
		userID int,
		amount float64,
	) error

	ValidateInvoice(
		ctx context.Context,
		userID int,
		invoiceID int,
		amount float64,
	) error

	MarkInvoicePaid(
		ctx context.Context,
		invoiceID int,
	) error
}

type TransactionUsecase interface {
	CreateTopUp(
		ctx context.Context,
		userID int,
		amount float64,
	) (*Transaction, error)

	CreatePayment(
		ctx context.Context,
		userID int,
		invoiceID int,
		amount float64,
	) (*Transaction, error)

	GetAll(
		ctx context.Context,
		userID int,
	) ([]Transaction, error)

	GetByID(
		ctx context.Context,
		id int,
	) (*Transaction, error)
}
