package repository

import (
	"context"
	"errors"

	"waste-management-service/internal/domain"

	"gorm.io/gorm"
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

	UpdateStatus(
		ctx context.Context,
		id int,
		status string,
	) error

	Delete(
		ctx context.Context,
		id int,
	) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) Create(
	ctx context.Context,
	transaction *domain.Transaction,
) error {
	return r.db.WithContext(ctx).
		Create(transaction).Error
}

func (r *transactionRepository) GetAllByUserID(
	ctx context.Context,
	userID int,
) ([]domain.Transaction, error) {

	var transactions []domain.Transaction

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("id ASC").
		Find(&transactions).Error

	return transactions, err
}

func (r *transactionRepository) GetByID(
	ctx context.Context,
	id int,
) (*domain.Transaction, error) {

	var transaction domain.Transaction

	err := r.db.WithContext(ctx).
		First(&transaction, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("transaction not found")
	}

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *transactionRepository) UpdateStatus(
	ctx context.Context,
	id int,
	status string,
) error {

	result := r.db.WithContext(ctx).
		Model(&domain.Transaction{}).
		Where("id = ?", id).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("transaction not found")
	}

	return nil
}

func (r *transactionRepository) Delete(
	ctx context.Context,
	id int,
) error {

	result := r.db.WithContext(ctx).
		Delete(&domain.Transaction{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("transaction not found")
	}

	return nil
}
