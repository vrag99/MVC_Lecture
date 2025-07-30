package main

import (
	"basic-requests/config"
	"basic-requests/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	_, err := config.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	router := mux.NewRouter()

	// Initialize UserController
	userController := controllers.NewUserController()

	// Define routes
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userController.GetUser).Methods("GET")
	router.HandleFunc("/users/add", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/update/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/delete/{id}", userController.DeleteUser).Methods("DELETE")

	fmt.Println("Server listening on http://localhost:8787")
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  /")
	fmt.Println("  GET  /users")
	fmt.Println("  GET  /users/{id}")
	fmt.Println("  POST /users/add")
	fmt.Println("  PUT  /users/update/{id}")
	fmt.Println("  DELETE /users/delete/{id}")

	serverErr := http.ListenAndServe(":8787", router)
	if serverErr != nil {
		fmt.Printf("Error starting server: %v\n", serverErr)
	}
}
