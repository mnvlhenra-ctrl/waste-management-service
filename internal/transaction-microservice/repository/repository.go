package repository

import (
	"context"

	"waste-management-service/internal/domain"
)

type TransactionRepository interface {
	Create(
		ctx context.Context,
		transaction *domain.Transaction,
	) error

	GetAllByUserID(
		ctx context.Context,
		userID int,
	) ([]domain.Transaction, error)

	GetByID(
		ctx context.Context,
		id int,
	) (*domain.Transaction, error)
}
