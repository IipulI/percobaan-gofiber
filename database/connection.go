package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB // Change this to a pointer

func Connection() error {
	urlConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", urlConn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(2 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)

	// Try to ping database
	if err := db.Ping(); err != nil {
		db.Close() // Close database connection on failure
		return fmt.Errorf("can't send ping to database: %w", err)
	}

	database = db // Correctly set the global database variable

	return nil
}

func GetDB() *sql.DB {
	return database
}
