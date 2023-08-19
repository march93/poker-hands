package validation

import (
	"log"
)

func Validate(input []string) {
	// Initialize ranks and suits to check against
	ranks := map[string]bool{"2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true, "T": true, "J": true, "Q": true, "K": true, "A": true}
	suits := map[string]bool{"S": true, "H": true, "D": true, "C": true}

	// Hand length should be 5
	if len(input) != 5 {
		log.Fatalf("Require 5 cards, %v cards given", len(input))
	}

	for _, card := range input {
		// Card length should be 2
		if len(card) != 2 {
			log.Fatalf("Invalid card format provided: %v", card)
		}

		// Card must belong to one of the ranks
		_, validRank := ranks[card[0:1]]
		if !validRank {
			log.Fatalf("%v - invalid rank provided: %v", card, card[0:1])
		}

		// Card must belong to one of the suits
		_, validSuit := suits[card[1:]]
		if !validSuit {
			log.Fatalf("%v - invalid suit provided: %v", card, card[1:])
		}
	}
}
