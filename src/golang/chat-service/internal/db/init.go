package db

import (
	"database/sql"
	"log"
)

func Init() (*sql.DB, error) {
	// Initialize the database connection and create necessary tables
	connStr := "host=localhost port=5433 user=chatservice dbname=chatservice sslmode=disable password=chatservice"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Test connection
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	return db, nil
}
