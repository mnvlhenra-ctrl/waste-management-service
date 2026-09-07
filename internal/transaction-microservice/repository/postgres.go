package repository

import (
	"context"
	"errors"

	"waste-management-service/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(
	db *pgxpool.Pool,
) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(
	ctx context.Context,
	transaction *domain.Transaction,
) error {

	query := `
		INSERT INTO transactions (
			user_id,
			invoice_id,
			name,
			amount,
			transaction_type,
			status,
			transaction_date
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	return r.db.QueryRow(
		ctx,
		query,
		transaction.UserID,
		transaction.InvoiceID,
		transaction.Name,
		transaction.Amount,
		transaction.TransactionType,
		transaction.Status,
		transaction.TransactionDate,
	).Scan(&transaction.ID)
}

func (r *PostgresRepository) GetAll(
	ctx context.Context,
) ([]domain.Transaction, error) {

	query := `
		SELECT
			id,
			user_id,
			invoice_id,
			name,
			amount,
			transaction_type,
			status,
			transaction_date
		FROM transactions
		ORDER BY id ASC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []domain.Transaction

	for rows.Next() {

		var transaction domain.Transaction

		err := rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.InvoiceID,
			&transaction.Name,
			&transaction.Amount,
			&transaction.TransactionType,
			&transaction.Status,
			&transaction.TransactionDate,
		)

		if err != nil {
			return nil, err
		}

		transactions = append(
			transactions,
			transaction,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *PostgresRepository) GetAllByUserID(
	ctx context.Context,
	userID int,
) ([]domain.Transaction, error) {

	query := `
		SELECT
			id,
			user_id,
			invoice_id,
			name,
			amount,
			transaction_type,
			status,
			transaction_date
		FROM transactions
		WHERE user_id = $1
		ORDER BY id ASC
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []domain.Transaction

	for rows.Next() {

		var transaction domain.Transaction

		err := rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.InvoiceID,
			&transaction.Name,
			&transaction.Amount,
			&transaction.TransactionType,
			&transaction.Status,
			&transaction.TransactionDate,
		)

		if err != nil {
			return nil, err
		}

		transactions = append(
			transactions,
			transaction,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *PostgresRepository) GetByID(
	ctx context.Context,
	id int,
) (*domain.Transaction, error) {

	query := `
		SELECT
			id,
			user_id,
			invoice_id,
			name,
			amount,
			transaction_type,
			status,
			transaction_date
		FROM transactions
		WHERE id = $1
	`

	var transaction domain.Transaction

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.InvoiceID,
		&transaction.Name,
		&transaction.Amount,
		&transaction.TransactionType,
		&transaction.Status,
		&transaction.TransactionDate,
	)

	if err != nil {
		return nil, errors.New("transaction not found")
	}

	return &transaction, nil
}

func (r *PostgresRepository) UpdateStatus(
	ctx context.Context,
	id int,
	status string,
) error {

	query := `
		UPDATE transactions
		SET status = $1
		WHERE id = $2
	`

	result, err := r.db.Exec(
		ctx,
		query,
		status,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("transaction not found")
	}

	return nil
}

func (r *PostgresRepository) Delete(
	ctx context.Context,
	id int,
) error {

	query := `
		DELETE FROM transactions
		WHERE id = $1
	`

	result, err := r.db.Exec(
		ctx,
		query,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("transaction not found")
	}

	return nil
}
