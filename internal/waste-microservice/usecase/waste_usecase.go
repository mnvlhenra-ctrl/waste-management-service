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

func NewWasteUsecase(wasteRepo domain.WasteRepository) *WasteUsecase {
	return &WasteUsecase{wasteRepo: wasteRepo}
}

// Fungsi helper untuk validasi kategori
func isValidCategory(category string) bool {
	validCategories := []string{"Organik", "Anorganik", "B3", "Residu"}
	for _, valid := range validCategories {
		if strings.EqualFold(category, valid) {
			return true
		}
	}
	return false
}

func (u *WasteUsecase) Create(ctx context.Context, waste *domain.Waste) error {
	if waste == nil {
		return errors.New("waste is required")
	}
	if strings.TrimSpace(waste.Name) == "" {
		return errors.New("waste name is required")
	}
	if waste.StockAvailability < 0 || waste.Costs < 0 {
		return errors.New("stock availability and costs cannot be negative")
	}
	if !isValidCategory(waste.Category) {
		return errors.New("invalid category. allowed: Organik, Anorganik, B3, Residu")
	}

	// Rapikan kapitalisasi kategori
	waste.Category = strings.Title(strings.ToLower(waste.Category))

	return u.wasteRepo.Create(ctx, waste)
}

func (u *WasteUsecase) GetAll(ctx context.Context, category string, isSeparated string) ([]domain.Waste, error) {
	return u.wasteRepo.GetAll(ctx, category, isSeparated)
}

func (u *WasteUsecase) GetByID(ctx context.Context, id int) (*domain.Waste, error) {
	if id <= 0 {
		return nil, errors.New("invalid waste id")
	}
	return u.wasteRepo.GetByID(ctx, id)
}

func (u *WasteUsecase) Update(ctx context.Context, waste *domain.Waste) error {
	if waste == nil || waste.ID <= 0 {
		return errors.New("invalid waste data")
	}
	if strings.TrimSpace(waste.Name) == "" {
		return errors.New("waste name is required")
	}
	if waste.StockAvailability < 0 || waste.Costs < 0 {
		return errors.New("stock availability and costs cannot be negative")
	}
	if !isValidCategory(waste.Category) {
		return errors.New("invalid category. allowed: Organik, Anorganik, B3, Residu")
	}

	waste.Category = strings.Title(strings.ToLower(waste.Category))

	return u.wasteRepo.Update(ctx, waste)
}

func (u *WasteUsecase) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid waste id")
	}
	return u.wasteRepo.Delete(ctx, id)
}
