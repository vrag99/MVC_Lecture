package main

import (
	"encoding/json"
	"fmt"
)

// 01 - basic types
func basic_types() {
	var name string = "John"
	var age int = 25
	var height float64 = 5.9
	var isVerified bool = true

	fmt.Println(name, age, height, isVerified)
}

// 02 - structs
type User struct {
	ID    int
	Name  string
	Email string
}

func structs() {
	u := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	fmt.Println(u)
}

// 03 - Json marshalling
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func json_marshalling() {
	p := Product{ID: 101, Name: "Pizza", Price: 9.99}
	jsonBytes, _ := json.Marshal(p)
	fmt.Println(string(jsonBytes))
}

// 04 - Json unmarshalling
type Order struct {
	UserID    int     `json:"user_id"`
	ProductID int     `json:"product_id"`
	Total     float64 `json:"total"`
}

func json_unmarshalling() {
	jsonData := `{"user_id": 1, "product_id": 101, "total": 19.98}`
	var o Order
	if err := json.Unmarshal([]byte(jsonData), &o); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%+v\n", o)
}

// 05 - Dynamic parsing
func dynamic_parsing() {
	jsonStr := `{"id": 1, "title": "Burger", "price": 5.99}`

	var data map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &data)

	fmt.Println("Title:", data["title"])
	fmt.Printf("Raw Map: %+v\n", data)
}

// Use this when struct type isn’t known.

func main() {
	basic_types()
	structs()
	json_marshalling()
	json_unmarshalling()
	dynamic_parsing()
}
