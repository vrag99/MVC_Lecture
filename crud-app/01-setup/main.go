package main

import (
	"crud-setup/config"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	config.InitDatabase()

	fmt.Println("Server listening on http://localhost:5000")
	http.ListenAndServe(":5000", nil)
}
