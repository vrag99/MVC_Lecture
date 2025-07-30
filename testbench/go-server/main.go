package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Server struct {
	db *sql.DB
}

type Request struct {
	UserID int    `json:"user_id"`
	Data   string `json:"data"`
}

type Response struct {
	ProcessedData  string    `json:"processed_data"`
	UserProfile    string    `json:"user_profile"`
	Timestamp      time.Time `json:"timestamp"`
	ProcessingTime string    `json:"processing_time"`
}

// CPU-intensive preprocessing: hash the data multiple times
func (s *Server) preprocessData(data string) string {
	hash := data
	// Simulate CPU-intensive work by hashing 1000 times
	for i := 0; i < 1000; i++ {
		hasher := sha256.New()
		hasher.Write([]byte(hash))
		hash = hex.EncodeToString(hasher.Sum(nil))
	}
	return hash[:16] // Return first 16 chars
}

// Database I/O operation
func (s *Server) getUserProfile(userID int) (string, error) {
	var profile string
	query := "SELECT profile_data FROM users WHERE id = ?"
	err := s.db.QueryRow(query, userID).Scan(&profile)
	if err != nil {
		if err == sql.ErrNoRows {
			// Create a mock profile if user doesn't exist
			profile = fmt.Sprintf("mock_profile_user_%d", userID)
		} else {
			return "", err
		}
	}
	return profile, nil
}

// CPU-intensive postprocessing: simulate complex calculations
func (s *Server) postprocessData(preprocessed string, profile string) string {
	combined := preprocessed + profile

	// Simulate CPU-intensive calculations
	result := 0
	for i := 0; i < 100000; i++ {
		result += int(combined[i%len(combined)]) * rand.Intn(10)
	}

	// More hashing for additional CPU work
	hasher := sha256.New()
	hasher.Write([]byte(combined + strconv.Itoa(result)))
	finalHash := hex.EncodeToString(hasher.Sum(nil))

	return finalHash[:32]
}

func (s *Server) processHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Step 1: CPU-bound preprocessing
	preprocessed := s.preprocessData(req.Data)

	// Step 2: DB I/O
	profile, err := s.getUserProfile(req.UserID)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Step 3: CPU-bound postprocessing
	finalResult := s.postprocessData(preprocessed, profile)

	processingTime := time.Since(startTime)

	response := Response{
		ProcessedData:  finalResult,
		UserProfile:    profile,
		Timestamp:      time.Now(),
		ProcessingTime: processingTime.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func initDB() (*sql.DB, error) {
	// Get database configuration from environment variables
	dbHost := getEnv("DB_HOST", "localhost")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "abcd")
	dbName := getEnv("DB_NAME", "performance_test")
	
	// Connection string format: username:password@tcp(host:port)/dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)

	var db *sql.DB
	var err error

	// Retry connection logic
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Failed to open database connection (attempt %d/%d): %v", i+1, maxRetries, err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Test the connection
		if err = db.Ping(); err != nil {
			log.Printf("Failed to ping database (attempt %d/%d): %v", i+1, maxRetries, err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Connection successful
		break
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %v", maxRetries, err)
	}

	// Create table if it doesn't exist
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INT PRIMARY KEY,
		profile_data VARCHAR(255) NOT NULL
	)`

	if _, err := db.Exec(createTable); err != nil {
		return nil, err
	}

	// Insert some sample data
	sampleData := []struct {
		id      int
		profile string
	}{
		{1, "user1_profile_data_with_some_content"},
		{2, "user2_profile_data_with_different_content"},
		{3, "user3_profile_data_with_more_content"},
		{4, "user4_profile_data_with_additional_content"},
		{5, "user5_profile_data_with_extended_content"},
	}

	for _, data := range sampleData {
		_, err := db.Exec("INSERT IGNORE INTO users (id, profile_data) VALUES (?, ?)",
			data.id, data.profile)
		if err != nil {
			log.Printf("Error inserting sample data: %v", err)
		}
	}

	return db, nil
}

func main() {
	// Initialize database
	db, err := initDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Set connection pool settings for better performance
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	server := &Server{db: db}

	// Set up routes
	r := mux.NewRouter()
	r.HandleFunc("/process", server.processHandler).Methods("POST")
	r.HandleFunc("/health", server.healthHandler).Methods("GET")

	// Configure HTTP server
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Go server starting on :8080")
	log.Fatal(srv.ListenAndServe())
}
