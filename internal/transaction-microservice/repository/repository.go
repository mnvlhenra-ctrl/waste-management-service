package repository

import (
	"context"
	"errors"

	"waste-management-service/internal/domain"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) domain.TransactionRepository {
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

func (r *transactionRepository) UpdateUserBalance(
	ctx context.Context,
	userID int,
	amount float64,
) error {

	result := r.db.WithContext(ctx).
		Table("users").
		Where("id = ?", userID).
		UpdateColumn("balance", gorm.Expr("balance + ?", amount))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *transactionRepository) ValidateInvoice(
	ctx context.Context,
	userID int,
	invoiceID int,
	amount float64,
) error {

	var invoice struct {
		ID     int
		Amount float64
		Status string
		House  struct {
			UserID int
		} `gorm:"foreignKey:HouseID"`
		HouseID int
	}

	err := r.db.WithContext(ctx).
		Table("invoices").
		Where("invoices.id = ?", invoiceID).
		First(&invoice).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("invoice not found")
	}

	if err != nil {
		return err
	}

	if invoice.Status != "UNPAID" {
		return errors.New("invoice has already been paid")
	}

	if invoice.Amount != amount {
		return errors.New("payment amount does not match invoice amount")
	}

	var house struct {
		UserID int
	}

	err = r.db.WithContext(ctx).
		Table("houses").
		Where("id = ?", invoice.HouseID).
		First(&house).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("house not found")
	}

	if err != nil {
		return err
	}

	if house.UserID != userID {
		return errors.New("invoice does not belong to user")
	}

	return nil
}

func (r *transactionRepository) MarkInvoicePaid(
	ctx context.Context,
	invoiceID int,
) error {

	result := r.db.WithContext(ctx).
		Table("invoices").
		Where("id = ?", invoiceID).
		Update("status", "PAID")

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("invoice not found")
	}

	return nil
}
