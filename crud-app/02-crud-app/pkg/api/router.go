package api

import (
	"crud-app/pkg/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouter configures and returns a new router with all API routes
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Initialize controllers
	userController := controllers.NewUserController()

	// Define routes
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userController.GetUser).Methods("GET")
	router.HandleFunc("/users/add", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/update/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/delete/{id}", userController.DeleteUser).Methods("DELETE")

	return router
}

// homeHandler handles the root endpoint
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// PrintRoutes prints all available routes for debugging
func PrintRoutes() {
	fmt.Println("Server listening on http://localhost:8787")
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  /")
	fmt.Println("  GET  /users")
	fmt.Println("  GET  /users/{id}")
	fmt.Println("  POST /users/add")
	fmt.Println("  PUT  /users/update/{id}")
	fmt.Println("  DELETE /users/delete/{id}")
}
