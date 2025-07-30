package models

import (
	"database/sql"
	"fmt"
)

// User represents a user entity
type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Country string `json:"country"`
}

// GetAllUsers retrieves all users from database
func GetAllUsers() ([]User, error) {
	query := "SELECT id, name, address, country FROM users"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying users: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Address, &user.Country)
		if err != nil {
			return nil, fmt.Errorf("error scanning user: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(id int) (*User, error) {
	query := "SELECT id, name, address, country FROM users WHERE id = ?"
	row := DB.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Address, &user.Country)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error scanning user: %v", err)
	}

	return &user, nil
}

// CreateUser creates a new user
func CreateUser(user User) (int, error) {
	query := "INSERT INTO users (name, address, country) VALUES (?, ?, ?)"
	result, err := DB.Exec(query, user.Name, user.Address, user.Country)
	if err != nil {
		return 0, fmt.Errorf("error creating user: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert id: %v", err)
	}

	return int(id), nil
}

// UpdateUser updates an existing user
func UpdateUser(id int, user User) error {
	query := "UPDATE users SET name = ?, address = ?, country = ? WHERE id = ?"
	result, err := DB.Exec(query, user.Name, user.Address, user.Country, id)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// DeleteUser deletes a user by ID
func DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
