package repositories

import (
	"context"
	"dealer_golang/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FavoriteRepository struct {
	DB *pgxpool.Pool
}

func NewFavoriteRepository(db *pgxpool.Pool) *FavoriteRepository {
	return &FavoriteRepository{DB: db}
}

func (r *FavoriteRepository) Add(userID int, vehicleID int) error {
	_, err := r.DB.Exec(
		context.Background(),
		`INSERT INTO favorites (user_id, vehicle_id) VALUES ($1, $2)`,
		userID, vehicleID,
	)
	return err
}

func (r *FavoriteRepository) Remove(userID int, vehicleID int) error {
	_, err := r.DB.Exec(
		context.Background(),
		`DELETE FROM favorites WHERE user_id = $1 AND vehicle_id = $2`,
		userID, vehicleID,
	)
	return err
}

func (r *FavoriteRepository) GetByUser(userID int) ([]models.Vehicle, error) {
	ctx := context.Background()

	query := `
		SELECT 
			v.vehicle_id,
			v.type_id,
			v.brand_id,
			v.name,
			v.fuel_type,
			v.transmission,
			v.price,
			v.stock,
			v.created_at,
			v.updated_at,
			t.type_name,
			b.brand_name
		FROM favorites f
		JOIN vehicles v ON f.vehicle_id = v.vehicle_id
		LEFT JOIN vehicle_types t ON t.type_id = v.type_id
		LEFT JOIN brands b ON b.brand_id = v.brand_id
		WHERE f.user_id = $1
	`

	rows, err := r.DB.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Vehicle

	for rows.Next() {
		var v models.Vehicle

		err := rows.Scan(
			&v.ID,
			&v.TypeID,
			&v.BrandID,
			&v.Name,
			&v.FuelType,
			&v.Transmission,
			&v.Price,
			&v.Stock,
			&v.CreatedAt,
			&v.UpdatedAt,
			&v.VehicleType,
			&v.Brand,
		)

		if err != nil {
			return nil, err
		}

		list = append(list, v)
	}

	return list, rows.Err()
}
