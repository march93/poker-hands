package validation

import (
	"fmt"
)

// Initialize ranks and suits to check against

// Integer representation of ranks
var Ranks = map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}

// Mapping of suits
var Suits = map[string]string{"S": "Spades", "H": "Hearts", "D": "Diamonds", "C": "Clubs"}

func Validate(input []string) error {
	// Hand length should be 5
	if len(input) != 5 {
		return fmt.Errorf("require 5 cards, %v card(s) given", len(input))
	}

	for _, card := range input {
		// Card length should be 2
		if len(card) != 2 {
			return fmt.Errorf("invalid card format provided: %v", card)
		}

		// Card must belong to one of the ranks
		_, validRank := Ranks[card[0:1]]
		if !validRank {
			return fmt.Errorf("%v - invalid rank provided: %c", card, card[0])
		}

		// Card must belong to one of the suits
		_, validSuit := Suits[card[1:]]
		if !validSuit {
			return fmt.Errorf("%v - invalid suit provided: %c", card, card[1])
		}
	}

	return nil
}
