package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDatabase() error {
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
		return fmt.Errorf("error opening database: %v", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	fmt.Println("hey db is connected!")
	return nil
}
