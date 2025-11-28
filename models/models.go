package models

import "time"

type User struct {
	ID        int       `db:"user_id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Vehicle struct {
	ID            int       `db:"vehicle_id"`
	TypeID        int       `db:"type_id"`         // FK ke vehicle_types
	BrandID       int       `db:"brand_id"`        // FK ke brands
	Name          string    `db:"name"`
	FuelType      string    `db:"fuel_type"`
	Transmission  string    `db:"transmission"`
	Price         float64   `db:"price"`
	Stock         int       `db:"stock"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`

	// hasil JOIN (opsional)
	VehicleType      string    `db:"type_name"`
	Brand     		 string    `db:"brand_name"`
}


type Transaction struct {
	ID          int       `db:"transaction_id"`
	UserID      int       `db:"user_id"`
	TotalAmount float64   `db:"total_amount"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type TransactionItem struct {
	ID            int     `db:"detail_id"`
	TransactionID int     `db:"transaction_id"`
	VehicleID     int     `db:"vehicle_id"`
	Quantity      int     `db:"quantity"`
	Price         float64 `db:"price"`
}

type PaymentDetail struct {
	ID            int        `db:"payment_id"`
	TransactionID int        `db:"transaction_id"`
	PaymentMethod string     `db:"payment_method"`
	Status        string     `db:"status"`
	PaidAt        *time.Time `db:"paid_at"` // null allowed
	Note		  *string     `db:"note"`
}

type Favorite struct {
	UserID    int       `db:"user_id"`
	VehicleID int       `db:"vehicle_id"`
	CreatedAt time.Time `db:"created_at"`
}

// Struct untuk menampung hasil JOIN transaction_items + vehicles
type TransactionItemWithVehicle struct {
	ID            int     `db:"detail_id"`
	TransactionID int     `db:"transaction_id"`
	VehicleID     int     `db:"vehicle_id"`
	Quantity      int     `db:"quantity"`
	Price         float64 `db:"price"`

	Brand string `db:"brand_name"`
	VehicleType  string `db:"type_name"`
	Name      string `db:"name"`
}


type TransactionWithPayment struct {
	ID            int
	UserID        int
	UserName      string
	TotalAmount   float64
	CreatedAt     time.Time
	PaymentMethod string
	Status        string
}

type VehicleType struct {
	ID       int    `db:"type_id"`
	TypeName string `db:"type_name"`
}

type Brand struct {
	ID        int    `db:"brand_id"`
	BrandName string `db:"brand_name"`
}

type VehicleFavoriteReport struct {
    ID             int     `db:"vehicle_id"`
    Name           string  `db:"name"`
    Price          float64 `db:"price"`
    Brand          string  `db:"brand_name"`
    Type           string  `db:"type_name"`
    TotalFavorites int     `db:"total_favorites"`
}
