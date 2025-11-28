package repositories

import (
	"context"
	"dealer_golang/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"fmt"
)

type TypeRepository struct {
	DB *pgxpool.Pool
}

func NewTypeRepository(db *pgxpool.Pool) *TypeRepository {
	return &TypeRepository{DB: db}
}

func (r *TypeRepository) GetAllTypes() ([]models.VehicleType, error) {
	query := `SELECT type_id, type_name FROM vehicle_types ORDER BY type_id`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.VehicleType

	for rows.Next() {
		var vt models.VehicleType
		err := rows.Scan(&vt.ID, &vt.TypeName)
		if err != nil {
			return nil, err
		}
		list = append(list, vt)
	}

	return list, rows.Err()
}

func (r *TypeRepository) CreateType(name string) error {
    exists, err := r.ExistsByName(name)
    if err != nil {
        return err
    }
    if exists {
        return fmt.Errorf("tipe '%s' sudah ada", name)
    }

    _, err = r.DB.Exec(
        context.Background(),
        `INSERT INTO vehicle_types (type_name) VALUES ($1)`,
        name,
    )
    return err
}


func (r *TypeRepository) ExistsByName(name string) (bool, error) {
    var exists bool
    err := r.DB.QueryRow(
        context.Background(),
        `SELECT EXISTS(SELECT 1 FROM vehicle_types WHERE LOWER(type_name) = LOWER($1))`,
        name,
    ).Scan(&exists)
    return exists, err
}

func (r *TypeRepository) GetTypeByName(name string) (models.VehicleType, error) {
    var vt models.VehicleType

    err := r.DB.QueryRow(
        context.Background(),
        `SELECT type_id, type_name
         FROM vehicle_types
         WHERE LOWER(type_name) = LOWER($1)
         LIMIT 1`,
        name,
    ).Scan(&vt.ID, &vt.TypeName)

    return vt, err
}



