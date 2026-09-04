package usecase

import (
	"context"
	"errors"
	"time"
	"waste-management-service/internal/domain"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo  domain.UserRepository
	jwtSecret string
}

func NewUserUsecase(repo domain.UserRepository, secret string) domain.UserUsecase {
	return &userUsecase{
		userRepo:  repo,
		jwtSecret: secret,
	}
}

func (u *userUsecase) Register(ctx context.Context, user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return u.userRepo.Store(ctx, user)
}

func (u *userUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString([]byte(u.jwtSecret))
}

func (u *userUsecase) GetProfile(ctx context.Context, id int) (domain.User, error) {
	return u.userRepo.GetByID(ctx, id)
}
