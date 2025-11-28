package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("failed to parse DATABASE_URL:", err)
	}

	cfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	DB, err = pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Fatal("failed to create connection pool:", err)
	}

	// Test koneksi
	if err := DB.Ping(context.Background()); err != nil {
		log.Fatal("failed to ping database:", err)
	}
}
