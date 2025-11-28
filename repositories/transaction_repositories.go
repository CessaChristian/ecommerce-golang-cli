package repositories

import (
	"context"
	"dealer_golang/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository struct {
	DB *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

// Insert Transaction Header
func (r *TransactionRepository) Create(t models.Transaction) (int, error) {
	query := `
		INSERT INTO transactions (user_id, total_amount)
		VALUES ($1, $2)
		RETURNING transaction_id
	`
	var id int
	err := r.DB.QueryRow(context.Background(), query,
		t.UserID, t.TotalAmount,
	).Scan(&id)

	return id, err
}

// Insert Transaction Item
func (r *TransactionRepository) AddItem(item models.TransactionItem) error {
	query := `
		INSERT INTO transaction_items (transaction_id, vehicle_id, quantity, price)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.DB.Exec(context.Background(), query,
		item.TransactionID,
		item.VehicleID,
		item.Quantity,
		item.Price,
	)
	return err
}

// Get All transactions by user
func (r *TransactionRepository) GetByUser(userID int) ([]models.Transaction, error) {
	ctx := context.Background()

	query := `
		SELECT transaction_id, user_id, total_amount, created_at, updated_at
		FROM transactions
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.DB.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Transaction

	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(
			&t.ID,
			&t.UserID,
			&t.TotalAmount,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, t)
	}

	return list, rows.Err()
}

// Get Transaction header by ID
func (r *TransactionRepository) GetTransactionByID(id int) (models.Transaction, error) {
	ctx := context.Background()

	var t models.Transaction

	err := r.DB.QueryRow(ctx, `
		SELECT transaction_id, user_id, total_amount, created_at, updated_at
		FROM transactions
		WHERE transaction_id = $1
	`, id).Scan(
		&t.ID,
		&t.UserID,
		&t.TotalAmount,
		&t.CreatedAt,
		&t.UpdatedAt,
	)

	return t, err
}

// Get all items for a transaction, joined with vehicle info
func (r *TransactionRepository) GetTransactionItems(transactionID int) ([]models.TransactionItemWithVehicle, error) {
	ctx := context.Background()

	query := `
		SELECT 
			ti.detail_id,
			ti.transaction_id,
			ti.vehicle_id,
			ti.quantity,
			ti.price,
			b.brand_name,
			t.type_name,
			v.name
		FROM transaction_items ti
		JOIN vehicles v ON v.vehicle_id = ti.vehicle_id
		JOIN brands b ON b.brand_id = v.brand_id
		JOIN vehicle_types t ON t.type_id = v.type_id
		WHERE ti.transaction_id = $1
	`

	rows, err := r.DB.Query(ctx, query, transactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.TransactionItemWithVehicle

	for rows.Next() {
		var it models.TransactionItemWithVehicle
		err := rows.Scan(
			&it.ID,
			&it.TransactionID,
			&it.VehicleID,
			&it.Quantity,
			&it.Price,
			&it.Brand,
			&it.VehicleType,
			&it.Name,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, it)
	}

	return items, rows.Err()
}

// Get ALL transactions + payment summary (admin)
func (r *TransactionRepository) GetAllTransactions() ([]models.TransactionWithPayment, error) {
	ctx := context.Background()

	query := `
		SELECT 
			t.transaction_id,
			t.user_id,
			u.name,
			t.total_amount,
			t.created_at,
			pd.payment_method,
			pd.status
		FROM transactions t
		JOIN users u ON u.user_id = t.user_id
		JOIN payment_detail pd ON pd.transaction_id = t.transaction_id
		ORDER BY t.created_at DESC
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.TransactionWithPayment

	for rows.Next() {
		var tr models.TransactionWithPayment
		err := rows.Scan(
			&tr.ID,
			&tr.UserID,
			&tr.UserName,
			&tr.TotalAmount,
			&tr.CreatedAt,
			&tr.PaymentMethod,
			&tr.Status,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, tr)
	}

	return list, rows.Err()
}
