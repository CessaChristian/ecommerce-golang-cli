package repositories

import (
	"context"
	"dealer_golang/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type VehicleRepository struct {
	DB *pgxpool.Pool
}

func NewVehicleRepository(db *pgxpool.Pool) *VehicleRepository {
	return &VehicleRepository{DB: db}
}

func (r *VehicleRepository) Create(v models.Vehicle) error {
	query := `
		INSERT INTO vehicles (type_id, brand_id, name, fuel_type, transmission, price, stock)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.DB.Exec(context.Background(), query,
		v.TypeID, v.BrandID, v.Name,
		v.FuelType, v.Transmission, v.Price, v.Stock,
	)
	return err
}

func (r *VehicleRepository) GetAll() ([]models.Vehicle, error) {
	query := `
		SELECT 
			v.vehicle_id, v.type_id, v.brand_id, v.name,
			v.fuel_type, v.transmission, v.price, v.stock,
			v.created_at, v.updated_at,
			t.type_name, b.brand_name
		FROM vehicles v
		LEFT JOIN vehicle_types t ON t.type_id = v.type_id
		LEFT JOIN brands b ON b.brand_id = v.brand_id
		ORDER BY v.vehicle_id
	`

	rows, err := r.DB.Query(context.Background(), query)
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

	return list, nil
}


func (r *VehicleRepository) GetByID(id int) (models.Vehicle, error) {
	query := `
		SELECT 
			v.vehicle_id, v.type_id, v.brand_id, v.name,
			v.fuel_type, v.transmission, v.price, v.stock,
			v.created_at, v.updated_at,
			t.type_name, b.brand_name
		FROM vehicles v
		LEFT JOIN vehicle_types t ON t.type_id = v.type_id
		LEFT JOIN brands b ON b.brand_id = v.brand_id
		WHERE v.vehicle_id = $1
	`

	var v models.Vehicle

	err := r.DB.QueryRow(context.Background(), query, id).Scan(
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

	return v, err
}


func (r *VehicleRepository) GetLowStock() ([]models.Vehicle, error) {
	query := `
		SELECT 
			v.vehicle_id, v.type_id, v.brand_id, v.name,
			v.fuel_type, v.transmission, v.price, v.stock,
			v.created_at, v.updated_at,
			t.type_name, b.brand_name
		FROM vehicles v
		LEFT JOIN vehicle_types t ON t.type_id = v.type_id
		LEFT JOIN brands b ON b.brand_id = v.brand_id
		WHERE v.stock <= 3
		ORDER BY v.stock ASC
	`

	rows, err := r.DB.Query(context.Background(), query)
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

	return list, nil
}


func (r *VehicleRepository) UpdateStock(id int, newStock int) error {
	query := `
		UPDATE vehicles SET stock = $1, updated_at = NOW()
		WHERE vehicle_id = $2
	`
	_, err := r.DB.Exec(context.Background(), query, newStock, id)
	return err
}

func (r *VehicleRepository) GetMostFavorited() ([]models.VehicleFavoriteReport, error) {
    query := `
        SELECT 
            v.vehicle_id,
            v.name,
            v.price,
            b.brand_name,
            t.type_name,
            COUNT(f.user_id) AS total_favorites
        FROM vehicles v
        LEFT JOIN favorites f ON f.vehicle_id = v.vehicle_id
        LEFT JOIN brands b ON b.brand_id = v.brand_id
        LEFT JOIN vehicle_types t ON t.type_id = v.type_id
        GROUP BY v.vehicle_id, b.brand_name, t.type_name
        ORDER BY total_favorites DESC;
    `

    rows, err := r.DB.Query(context.Background(), query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []models.VehicleFavoriteReport

    for rows.Next() {
        var d models.VehicleFavoriteReport
        err := rows.Scan(
            &d.ID,
            &d.Name,
            &d.Price,
            &d.Brand,
            &d.Type,
            &d.TotalFavorites,
        )
        if err != nil {
            return nil, err
        }
        list = append(list, d)
    }

    return list, nil
}

