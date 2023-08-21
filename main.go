package main

import (
	"fmt"
	"poker-hands/utils"
	"poker-hands/validation"
	"strings"
)

func main() {
	fmt.Println("Input one or more poker hands for evaluation.")
	fmt.Println("Each hand should be a comma-delimited string of 5 cards in the following format: `XX,XX,XX,XX,XX`.")
	fmt.Println("The first character represents the rank from 2-9, T, J, Q, K, A.")
	fmt.Println("The second character represents the suit as S (Spades), H (Hearts), D (Diamonds), C (Clubs).")
	fmt.Println("e.g. 2H,3S,TD,6C,JH")
	fmt.Println()

	fmt.Println("Enter poker hand(s) and click return to enter more.")
	fmt.Println("Enter `done` to start evaluating.")
	fmt.Println()

	var input string
	fmt.Scanln(&input)
	hands := [][]string{}

	for input != "done" {
		s := strings.Split(input, ",")

		// Validate input for hand length, card value
		validation.Validate(s)

		// Add valid input to hands slice
		hands = append(hands, s)

		// Continue scanning for more input
		input = ""
		fmt.Scanln(&input)
	}

	fmt.Printf("\nAll inputs: %v\n\n", hands)
	evaluate(hands)
}

func evaluate(input [][]string) {
	for _, hand := range input {
		// Transform into card structs
		cards := utils.Transform(hand)
		fmt.Printf("\n\nTransformed cards: %v\n\n", cards)
	}
}
