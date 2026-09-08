package usecase

import (
	"context"
	"errors"
	"time"

	"waste-management-service/internal/domain"
)

const (
	TransactionTypeTopUp   = "TOP_UP"
	TransactionTypePayment = "PAYMENT"

	TransactionStatusSuccess = "SUCCESS"
	TransactionStatusFailed  = "FAILED"
)

type TransactionUsecase struct {
	transactionRepo domain.TransactionRepository
}

func NewTransactionUsecase(
	transactionRepo domain.TransactionRepository,
) *TransactionUsecase {
	return &TransactionUsecase{
		transactionRepo: transactionRepo,
	}
}

// CreateTopUp creates a TOP_UP transaction.
func (u *TransactionUsecase) CreateTopUp(
	ctx context.Context,
	userID int,
	amount float64,
) (*domain.Transaction, error) {

	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	transaction := &domain.Transaction{
		UserID:          userID,
		Name:            "Top Up Saldo",
		Amount:          amount,
		TransactionType: TransactionTypeTopUp,
		Status:          TransactionStatusSuccess,
		TransactionDate: time.Now(),
	}

	if err := u.transactionRepo.Create(
		ctx,
		transaction,
	); err != nil {
		return nil, err
	}

	if err := u.transactionRepo.UpdateUserBalance(
		ctx,
		userID,
		amount,
	); err != nil {
		return nil, err
	}

	return transaction, nil
}

// CreatePayment creates a PAYMENT transaction.
func (u *TransactionUsecase) CreatePayment(
	ctx context.Context,
	userID int,
	invoiceID int,
	amount float64,
) (*domain.Transaction, error) {

	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	if invoiceID <= 0 {
		return nil, errors.New("invalid invoice id")
	}

	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	if err := u.transactionRepo.ValidateInvoice(
		ctx,
		userID,
		invoiceID,
		amount,
	); err != nil {
		return nil, err
	}

	// Kurangi saldo user sesuai jumlah pembayaran
	if err := u.transactionRepo.UpdateUserBalance(
		ctx,
		userID,
		-amount,
	); err != nil {
		return nil, err
	}

	transaction := &domain.Transaction{
		UserID:          userID,
		InvoiceID:       &invoiceID,
		Name:            "Pembayaran Tagihan",
		Amount:          amount,
		TransactionType: TransactionTypePayment,
		Status:          TransactionStatusSuccess,
		TransactionDate: time.Now(),
	}

	if err := u.transactionRepo.Create(
		ctx,
		transaction,
	); err != nil {
		return nil, err
	}

	if err := u.transactionRepo.MarkInvoicePaid(
		ctx,
		invoiceID,
	); err != nil {
		return nil, err
	}

	return transaction, nil
}

// GetAll returns all transactions for a user.
func (u *TransactionUsecase) GetAll(
	ctx context.Context,
	userID int,
) ([]domain.Transaction, error) {

	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	return u.transactionRepo.GetAllByUserID(
		ctx,
		userID,
	)
}

// GetByID returns a transaction by ID.
func (u *TransactionUsecase) GetByID(
	ctx context.Context,
	id int,
) (*domain.Transaction, error) {

	if id <= 0 {
		return nil, errors.New("invalid transaction id")
	}

	return u.transactionRepo.GetByID(
		ctx,
		id,
	)
}
