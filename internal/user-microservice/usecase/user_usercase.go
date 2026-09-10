package usecase

import (
	"context"
	"errors"
	"time"

	"waste-management-service/internal/domain"
	"waste-management-service/pkg/email"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo     domain.UserRepository
	houseRepo    domain.HouseRepository
	invoiceRepo  domain.InvoiceRepository
	emailService email.Service
	jwtSecret    string
}

func NewUserUsecase(
	repo domain.UserRepository,
	houseRepo domain.HouseRepository,
	invoiceRepo domain.InvoiceRepository,
	emailService email.Service,
	secret string,
) domain.UserUsecase {

	return &userUsecase{
		userRepo:     repo,
		houseRepo:    houseRepo,
		invoiceRepo:  invoiceRepo,
		emailService: emailService,
		jwtSecret:    secret,
	}
}

func (u *userUsecase) Register(
	ctx context.Context,
	user *domain.User,
	houseNumber string,
) error {

	if houseNumber == "" {
		return errors.New("house number is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	// Create user
	if err := u.userRepo.Store(ctx, user); err != nil {
		return err
	}

	// Create house automatically
	house := &domain.House{
		UserID:      user.ID,
		HouseNumber: houseNumber,
		Name:        "Rumah " + houseNumber,
		Category:    "Residential",
		Services:    "Senin: Organik, Kamis: Anorganik",
		Costs:       15000,
	}

	if err := u.houseRepo.Create(ctx, house); err != nil {
		return err
	}

	// Create invoice for current month automatically
	now := time.Now()

	billingMonth := now.Format("2006-01")

	// Invoice due date = end of current month
	nextMonth := time.Date(
		now.Year(),
		now.Month()+1,
		1,
		23,
		59,
		59,
		0,
		now.Location(),
	)

	dueDate := nextMonth.Add(-time.Second)

	invoice := &domain.Invoice{
		HouseID:      house.ID,
		BillingMonth: billingMonth,
		Amount:       house.Costs,
		Status:       "UNPAID",
		DueDate:      dueDate,
	}

	if err := u.invoiceRepo.Create(ctx, invoice); err != nil {
		return err
	}

	body := email.RegistrationEmail(
		user.Email,
		user.Email,
	)

	if err := u.emailService.Send(
		ctx,
		user.Email,
		"Registration Successful",
		body,
	); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) Login(
	ctx context.Context,
	email,
	password string,
) (string, error) {

	user, err := u.userRepo.GetByEmail(ctx, email)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
		},
	)

	return token.SignedString([]byte(u.jwtSecret))
}

func (u *userUsecase) GetProfile(
	ctx context.Context,
	id int) (domain.User, error) {

	return u.userRepo.GetByID(ctx, id)
}
