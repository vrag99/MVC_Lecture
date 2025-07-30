package models

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

// * NOTE: When you actually do your assignment, make utility loaders and a config package for env references.

func InitDatabase() (*sql.DB, error) {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found")
	}

	// Get database connection details from environment
	dbHost := os.Getenv("MYSQL_HOST")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	dbPort := os.Getenv("MYSQL_PORT")

	// Create connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, database)

	// Open database connection
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Configure connection pool settings
	DB.SetMaxOpenConns(25)                 // Maximum number of open connections
	DB.SetMaxIdleConns(5)                  // Maximum number of idle connections
	DB.SetConnMaxLifetime(5 * time.Minute) // Maximum lifetime of a connection

	// Test the connection
	err = DB.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	fmt.Println("Database connected successfully!")
	return DB, nil
}

// CloseDatabase closes the database connection
func CloseDatabase() error {
	if DB != nil {
		fmt.Println("Closing database connection...")
		return DB.Close()
	}
	return nil
}
