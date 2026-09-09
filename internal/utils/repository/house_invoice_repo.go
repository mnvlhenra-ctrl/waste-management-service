package repository

import (
	"context"
	"errors"

	"waste-management-service/internal/domain"

	"gorm.io/gorm"
)

// ====================
// HOUSE REPOSITORY
// ====================

type houseRepository struct {
	db *gorm.DB
}

func NewHouseRepository(db *gorm.DB) domain.HouseRepository {
	return &houseRepository{
		db: db,
	}
}

func (r *houseRepository) Create(
	ctx context.Context,
	house *domain.House,
) error {
	return r.db.WithContext(ctx).
		Create(house).Error
}

// ====================
// INVOICE REPOSITORY
// ====================

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) domain.InvoiceRepository {
	return &invoiceRepository{
		db: db,
	}
}

func (r *invoiceRepository) Create(
	ctx context.Context,
	invoice *domain.Invoice,
) error {
	return r.db.WithContext(ctx).
		Create(invoice).Error
}

func (r *invoiceRepository) GetByID(
	ctx context.Context,
	id int,
) (*domain.Invoice, error) {

	var invoice domain.Invoice

	err := r.db.WithContext(ctx).
		First(&invoice, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("invoice not found")
	}

	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

// GetByUserID mengambil semua invoice milik user
func (r *invoiceRepository) GetByUserID(
	ctx context.Context,
	userID int,
) ([]domain.Invoice, error) {

	var invoices []domain.Invoice

	err := r.db.WithContext(ctx).
		Table("invoices").
		Joins("JOIN houses ON houses.id = invoices.house_id").
		Where("houses.user_id = ?", userID).
		Order("invoices.id ASC").
		Find(&invoices).Error

	return invoices, err
}

func (r *invoiceRepository) UpdateStatus(
	ctx context.Context,
	id int,
	status string,
) error {

	result := r.db.WithContext(ctx).
		Model(&domain.Invoice{}).
		Where("id = ?", id).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("invoice not found")
	}

	return nil
}
