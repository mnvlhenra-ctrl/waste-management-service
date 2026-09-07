package usecase

import (
	"context"
	"errors"
	"strings"

	"waste-management-service/internal/domain"
)

type WasteUsecase struct {
	wasteRepo domain.WasteRepository
}

func NewWasteUsecase(
	wasteRepo domain.WasteRepository,
) *WasteUsecase {
	return &WasteUsecase{
		wasteRepo: wasteRepo,
	}
}

func (u *WasteUsecase) Create(
	ctx context.Context,
	waste *domain.Waste,
) error {

	if waste == nil {
		return errors.New("waste is required")
	}

	if strings.TrimSpace(waste.Name) == "" {
		return errors.New("waste name is required")
	}

	if waste.StockAvailability < 0 {
		return errors.New("stock availability cannot be negative")
	}

	if waste.Costs < 0 {
		return errors.New("costs cannot be negative")
	}

	if strings.TrimSpace(waste.Category) == "" {
		return errors.New("waste category is required")
	}

	return u.wasteRepo.Create(ctx, waste)
}

func (u *WasteUsecase) GetAll(
	ctx context.Context,
) ([]domain.Waste, error) {

	return u.wasteRepo.GetAll(ctx)
}

func (u *WasteUsecase) GetByID(
	ctx context.Context,
	id int,
) (*domain.Waste, error) {

	if id <= 0 {
		return nil, errors.New("invalid waste id")
	}

	return u.wasteRepo.GetByID(ctx, id)
}

func (u *WasteUsecase) Update(
	ctx context.Context,
	waste *domain.Waste,
) error {

	if waste == nil {
		return errors.New("waste is required")
	}

	if waste.ID <= 0 {
		return errors.New("invalid waste id")
	}

	if strings.TrimSpace(waste.Name) == "" {
		return errors.New("waste name is required")
	}

	if waste.StockAvailability < 0 {
		return errors.New("stock availability cannot be negative")
	}

	if waste.Costs < 0 {
		return errors.New("costs cannot be negative")
	}

	if strings.TrimSpace(waste.Category) == "" {
		return errors.New("waste category is required")
	}

	return u.wasteRepo.Update(ctx, waste)
}

func (u *WasteUsecase) Delete(
	ctx context.Context,
	id int,
) error {

	if id <= 0 {
		return errors.New("invalid waste id")
	}

	return u.wasteRepo.Delete(ctx, id)
}
