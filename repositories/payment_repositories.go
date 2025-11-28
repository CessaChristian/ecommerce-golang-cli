package repositories

import (
	"context"
	"dealer_golang/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentRepository struct {
	DB *pgxpool.Pool
}

func NewPaymentRepository(db *pgxpool.Pool) *PaymentRepository {
	return &PaymentRepository{DB: db}
}

// Create Payment Detail
func (r *PaymentRepository) Create(p models.PaymentDetail) error {
	query := `
		INSERT INTO payment_detail (transaction_id, payment_method, status, paid_at, note)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.DB.Exec(context.Background(), query,
		p.TransactionID,
		p.PaymentMethod,
		p.Status,
		p.PaidAt,
		p.Note,
	)

	return err
}

// Get Payment Detail (One-to-One)
func (r *PaymentRepository) GetByTransactionID(transactionID int) (models.PaymentDetail, error) {
	var p models.PaymentDetail

	query := `
		SELECT payment_id, transaction_id, payment_method, status, paid_at, note
		FROM payment_detail
		WHERE transaction_id = $1
	`

	err := r.DB.QueryRow(context.Background(), query, transactionID).Scan(
		&p.ID,
		&p.TransactionID,
		&p.PaymentMethod,
		&p.Status,
		&p.PaidAt,
		&p.Note,
	)

	return p, err
}

func (r *PaymentRepository) GetPaymentDetail(transactionID int) (models.PaymentDetail, error) {
    var p models.PaymentDetail

    query := `
        SELECT payment_id, transaction_id, payment_method, status, paid_at, note
        FROM payment_detail
        WHERE transaction_id = $1
    `

    err := r.DB.QueryRow(context.Background(), query, transactionID).Scan(
        &p.ID,
        &p.TransactionID,
        &p.PaymentMethod,
        &p.Status,
        &p.PaidAt,
        &p.Note,
    )

    return p, err
}

