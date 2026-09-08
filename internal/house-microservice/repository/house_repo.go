package repository

import (
	"context"

	"waste-management-service/internal/domain"

	"gorm.io/gorm"
)

type houseRepository struct {
	db *gorm.DB
}

func NewHouseRepository(db *gorm.DB) domain.HouseRepository {
	return &houseRepository{db}
}

func (r *houseRepository) Create(
	ctx context.Context,
	house *domain.House,
) error {
	return r.db.WithContext(ctx).Create(house).Error
}
