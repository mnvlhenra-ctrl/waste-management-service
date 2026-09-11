package repository

import (
	"context"

	"waste-management-service/internal/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Store(
	ctx context.Context,
	u *domain.User,
) error {
	return r.db.WithContext(ctx).
		Create(u).Error
}

func (r *userRepository) GetByEmail(
	ctx context.Context,
	email string,
) (domain.User, error) {

	var user domain.User

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error

	return user, err
}

func (r *userRepository) GetByID(
	ctx context.Context,
	id int,
) (domain.User, error) {

	var user domain.User

	err := r.db.WithContext(ctx).
		First(&user, id).Error

	return user, err
}
