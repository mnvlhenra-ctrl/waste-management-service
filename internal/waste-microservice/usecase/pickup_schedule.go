package usecase

import (
	"context"
	"errors"
	"time"
)

type PickupSchedule struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	PickupDate time.Time `json:"pickup_date"`
	WasteType  string    `json:"waste_type"`
}

type PickupScheduleRepository interface {
	Create(ctx context.Context, schedule *PickupSchedule) error
	GetTomorrow(ctx context.Context) ([]PickupSchedule, error)
}

type PickupScheduleUsecase struct {
	repo PickupScheduleRepository
}

func NewPickupScheduleUsecase(
	repo PickupScheduleRepository,
) *PickupScheduleUsecase {
	return &PickupScheduleUsecase{
		repo: repo,
	}
}

func (u *PickupScheduleUsecase) Create(
	ctx context.Context,
	schedule *PickupSchedule,
) error {

	if schedule == nil {
		return errors.New("pickup schedule is required")
	}

	if schedule.UserID <= 0 {
		return errors.New("invalid user id")
	}

	if schedule.PickupDate.IsZero() {
		return errors.New("pickup date is required")
	}

	if schedule.PickupDate.Before(time.Now()) {
		return errors.New("pickup date cannot be in the past")
	}

	if schedule.WasteType == "" {
		return errors.New("waste type is required")
	}

	return u.repo.Create(ctx, schedule)
}

func (u *PickupScheduleUsecase) GetTomorrow(
	ctx context.Context,
) ([]PickupSchedule, error) {

	return u.repo.GetTomorrow(ctx)
}
