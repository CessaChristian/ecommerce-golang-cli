package repositories

import (
	"context"
	"dealer_golang/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user models.User) error {
	query := `
		INSERT INTO users (name, email, password, role)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.DB.Exec(context.Background(), query,
		user.Name, user.Email, user.Password, user.Role,
	)
	return err
}

func (r *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User

	query := `
		SELECT user_id, name, email, password, role, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	err := r.DB.QueryRow(context.Background(), query, email).Scan(
		&user.ID, &user.Name, &user.Email,
		&user.Password, &user.Role,
		&user.CreatedAt, &user.UpdatedAt,
	)

	return user, err
}

func (r *UserRepository) GetByID(id int) (models.User, error) {
	var user models.User

	query := `
		SELECT user_id, name, email, password, role, created_at, updated_at
		FROM users
		WHERE user_id = $1
	`
	err := r.DB.QueryRow(context.Background(), query, id).Scan(
		&user.ID, &user.Name, &user.Email,
		&user.Password, &user.Role,
		&user.CreatedAt, &user.UpdatedAt,
	)
	return user, err
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	query := `
		SELECT user_id, name, email, role, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
	`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(
			&u.ID, &u.Name, &u.Email, &u.Role,
			&u.CreatedAt, &u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, u)
	}

	return list, nil
}
