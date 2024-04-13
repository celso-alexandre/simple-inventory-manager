package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
	// if DB != nil {
	// 	return DB, nil
	// }
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		return DB, errors.New("DB_URL not set")
	}
	fmt.Println("Connecting to DB", connStr)
	var err error
	DB, err = sql.Open("postgres", connStr)
	return DB, err
}
