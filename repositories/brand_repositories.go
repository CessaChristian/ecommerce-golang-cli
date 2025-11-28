package repositories

import (
	"context"
	"dealer_golang/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"fmt"
)

type BrandRepository struct {
	DB *pgxpool.Pool
}

func NewBrandRepository(db *pgxpool.Pool) *BrandRepository {
	return &BrandRepository{DB: db}
}

func (r *BrandRepository) GetAllBrands() ([]models.Brand, error) {
	query := `SELECT brand_id, brand_name FROM brands ORDER BY brand_id`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Brand

	for rows.Next() {
		var b models.Brand
		err := rows.Scan(&b.ID, &b.BrandName)
		if err != nil {
			return nil, err
		}
		list = append(list, b)
	}

	return list, rows.Err()
}

func (r *BrandRepository) CreateBrand(name string) error {
    exists, err := r.ExistsByName(name)
    if err != nil {
        return err
    }
    if exists {
        return fmt.Errorf("brand '%s' sudah ada", name)
    }

    _, err = r.DB.Exec(
        context.Background(),
        `INSERT INTO brands (brand_name) VALUES ($1)`,
        name,
    )
    return err
}

func (r *BrandRepository) ExistsByName(name string) (bool, error) {
    var exists bool
    err := r.DB.QueryRow(
        context.Background(),
        `SELECT EXISTS(SELECT 1 FROM brands WHERE LOWER(brand_name) = LOWER($1))`,
        name,
    ).Scan(&exists)
    return exists, err
}

func (r *BrandRepository) GetBrandByName(name string) (models.Brand, error) {
    var b models.Brand

    err := r.DB.QueryRow(
        context.Background(),
        `SELECT brand_id, brand_name
         FROM brands
         WHERE LOWER(brand_name) = LOWER($1)
         LIMIT 1`,
        name,
    ).Scan(&b.ID, &b.BrandName)

    return b, err
}

