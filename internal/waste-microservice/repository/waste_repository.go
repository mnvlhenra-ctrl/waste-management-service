package repository

import (
	"context"
	"errors"

	"waste-management-service/internal/domain"

	"gorm.io/gorm"
)

type wasteRepository struct {
	db *gorm.DB
}

func NewWasteRepository(db *gorm.DB) domain.WasteRepository {
	return &wasteRepository{
		db: db,
	}
}

func (r *wasteRepository) Create(
	ctx context.Context,
	waste *domain.Waste,
) error {
	return r.db.WithContext(ctx).
		Create(waste).
		Error
}

func (r *wasteRepository) GetAll(
	ctx context.Context,
	category string,
	isSeparated string,
) ([]domain.Waste, error) {

	var wastes []domain.Waste

	query := r.db.WithContext(ctx)

	// Filter category
	if category != "" {
		query = query.Where(
			"category = ?",
			category,
		)
	}

	// Filter separation status
	if isSeparated == "true" {
		query = query.Where(
			"is_separated = ?",
			true,
		)
	} else if isSeparated == "false" {
		query = query.Where(
			"is_separated = ?",
			false,
		)
	}

	err := query.
		Order("id ASC").
		Find(&wastes).
		Error

	return wastes, err
}

func (r *wasteRepository) GetByID(
	ctx context.Context,
	id int,
) (*domain.Waste, error) {

	var waste domain.Waste

	err := r.db.WithContext(ctx).
		First(&waste, id).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("waste not found")
	}

	return &waste, err
}

func (r *wasteRepository) Update(
	ctx context.Context,
	waste *domain.Waste,
) error {

	result := r.db.WithContext(ctx).
		Model(&domain.Waste{}).
		Where("id = ?", waste.ID).
		Updates(map[string]interface{}{
			"name":                waste.Name,
			"stock_availability":  waste.StockAvailability,
			"separated_costs":     waste.SeparatedCosts,
			"non_separated_costs": waste.NonSeparatedCosts,
			"category":            waste.Category,
			"is_separated":        waste.IsSeparated,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("waste not found")
	}

	return nil
}

func (r *wasteRepository) Delete(
	ctx context.Context,
	id int,
) error {

	result := r.db.WithContext(ctx).
		Delete(&domain.Waste{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("waste not found")
	}

	return nil
}
