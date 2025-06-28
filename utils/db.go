package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func DBConnect() (*pgxpool.Pool, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
	)

	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return pool, err
}
