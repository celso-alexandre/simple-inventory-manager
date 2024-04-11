package db

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
	if DB != nil {
		return DB, nil
	}
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		return nil, errors.New("DB_URL not set")
	}
	DB, err := sql.Open("postgres", connStr)
	return DB, err
}
