package domain

import "context"

type User struct {
	ID       int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Email    string  `json:"email" gorm:"unique;not null"`
	Password string  `json:"password" gorm:"not null"`
	Role     string  `json:"role" gorm:"default:'USER'"`
	Balance  float64 `json:"balance" gorm:"type:decimal(12,2);default:0"`
}

type UserRepository interface {
	Store(ctx context.Context, u *User) error
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, id int) (User, error)
	Delete(ctx context.Context, id int) error
}

type UserUsecase interface {
	Register(ctx context.Context, u *User, houseNumber string) error
	Login(ctx context.Context, email, password string) (string, error)
	GetProfile(ctx context.Context, id int) (User, error)
	DeleteUser(ctx context.Context, id int) error
}
