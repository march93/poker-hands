package main

import (
	"fmt"
	"os"
	"poker-hands/validation"
	"strings"
)

func main() {
	args := os.Args

	cards := strings.Split(args[1], ",")
	fmt.Println(cards)

	// Validate input for hand length, card value
	validation.Validate(cards)
}
