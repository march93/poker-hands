package main

import (
	"fmt"
	"os"
	"poker-hands/validation"
)

func main() {
	args := os.Args

	fmt.Println(args[1:])

	// Validate input for hand length, card value
	validation.Validate(args[1:])
}
