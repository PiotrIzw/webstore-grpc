package database

import (
	"database/sql"
	"fmt"
	"github.com/PiotrIzw/webstore-grcp/config/config"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func ConnectDB() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Construct the DSN (Data Source Name) for PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	// Open a connection to the database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Test the connection to ensure it's valid
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Assign the database connection to the global DB variable
	DB = db
	log.Println("Database connection established")
}
