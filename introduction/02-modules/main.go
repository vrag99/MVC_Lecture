package main

import (
	"fmt"
	"learn-modules/utils"
)

func main() {
	var name string

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	utils.Greet(name)
}
