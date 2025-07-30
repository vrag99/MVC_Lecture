package controllers

import (
	"crud-app/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// UserController handles user-related HTTP requests
type UserController struct{}

// NewUserController creates a new UserController
func NewUserController() *UserController {
	return &UserController{}
}

// GetUsers handles GET /users
func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching users: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUser handles GET /users/{id}
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := models.GetUserByID(id)
	if err != nil {
		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Error fetching user: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUser handles POST /users/add
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if user.Name == "" || user.Address == "" || user.Country == "" {
		http.Error(w, "Name, address, and country are required", http.StatusBadRequest)
		return
	}

	id, err := models.CreateUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message": "User created successfully",
		"id":      id,
	}
	json.NewEncoder(w).Encode(response)
}

// UpdateUser handles PUT /users/update/{id}
func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/users/update/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if user.Name == "" || user.Address == "" || user.Country == "" {
		http.Error(w, "Name, address, and country are required", http.StatusBadRequest)
		return
	}

	err = models.UpdateUser(id, user)
	if err != nil {
		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Error updating user: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": fmt.Sprintf("User with id %d updated successfully", id),
	}
	json.NewEncoder(w).Encode(response)
}

// DeleteUser handles DELETE /users/delete/{id}
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/users/delete/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = models.DeleteUser(id)
	if err != nil {
		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Error deleting user: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": fmt.Sprintf("User with id %d deleted successfully", id),
	}
	json.NewEncoder(w).Encode(response)
}
