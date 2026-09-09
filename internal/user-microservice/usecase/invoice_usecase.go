package usecase

import (
	"context"
	"errors"

	"waste-management-service/internal/domain"
)

type InvoiceUsecase struct {
	invoiceRepo domain.InvoiceRepository
}

func NewInvoiceUsecase(
	invoiceRepo domain.InvoiceRepository,
) *InvoiceUsecase {
	return &InvoiceUsecase{
		invoiceRepo: invoiceRepo,
	}
}

func (u *InvoiceUsecase) GetByUserID(
	ctx context.Context,
	userID int,
) ([]domain.Invoice, error) {

	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	return u.invoiceRepo.GetByUserID(ctx, userID)
}
