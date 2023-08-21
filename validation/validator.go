package validation

import (
	"log"
)

// Initialize ranks and suits to check against
var Ranks = map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}
var Suits = map[string]bool{"S": true, "H": true, "D": true, "C": true}

func Validate(input []string) {
	// Hand length should be 5
	if len(input) != 5 {
		log.Fatalf("Require 5 cards, %v card(s) given", len(input))
	}

	for _, card := range input {
		// Card length should be 2
		if len(card) != 2 {
			log.Fatalf("Invalid card format provided: %v", card)
		}

		// Card must belong to one of the ranks
		_, validRank := Ranks[card[0:1]]
		if !validRank {
			log.Fatalf("%v - invalid rank provided: %v", card, card[0])
		}

		// Card must belong to one of the suits
		_, validSuit := Suits[card[1:]]
		if !validSuit {
			log.Fatalf("%v - invalid suit provided: %v", card, card[1])
		}
	}
}
