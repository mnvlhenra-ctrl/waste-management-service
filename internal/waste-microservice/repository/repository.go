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
	return &wasteRepository{db}
}

func (r *wasteRepository) Create(
	ctx context.Context,
	waste *domain.Waste,
) error {
	return r.db.WithContext(ctx).Create(waste).Error
}

func (r *wasteRepository) GetAll(
	ctx context.Context,
) ([]domain.Waste, error) {

	var wastes []domain.Waste

	err := r.db.WithContext(ctx).
		Order("id ASC").
		Find(&wastes).Error

	return wastes, err
}

func (r *wasteRepository) GetByID(
	ctx context.Context,
	id int,
) (*domain.Waste, error) {

	var waste domain.Waste

	err := r.db.WithContext(ctx).
		First(&waste, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("waste not found")
	}

	if err != nil {
		return nil, err
	}

	return &waste, nil
}

func (r *wasteRepository) Update(
	ctx context.Context,
	waste *domain.Waste,
) error {

	result := r.db.WithContext(ctx).
		Model(&domain.Waste{}).
		Where("id = ?", waste.ID).
		Updates(map[string]interface{}{
			"name":               waste.Name,
			"stock_availability": waste.StockAvailability,
			"costs":              waste.Costs,
			"category":           waste.Category,
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
