package main

import (
	"context"
	"crud-app/pkg/api"
	"crud-app/pkg/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Initialize database
	_, err := models.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Setup router with all API routes
	router := api.SetupRouter()

	// Print available routes
	api.PrintRoutes()

	// Create HTTP server
	server := &http.Server{
		Addr:    ":8787",
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		fmt.Printf("Starting server on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	// Close database connection
	if err := models.CloseDatabase(); err != nil {
		log.Printf("Error closing database: %v", err)
	}

	fmt.Println("Server exited gracefully")
}
