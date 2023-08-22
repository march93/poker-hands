package utils

import (
	"poker-hands/validation"
)

/*	Integer representation of hand types
 *	Royal flush -> 9
 *	Straight flush -> 8
 *	Four of a kind -> 7
 *	Full house -> 6
 *	Flush -> 5
 *	Straight -> 4
 *	Three of a kind -> 3
 *	Two pair -> 2
 *	Pair -> 1
 *	High card -> 0
 */

type Hand struct {
	Cards []Card
	Type  int
	Score int
}

func determineHand() {}

// Function to determine if a hand is a royal flush
func isRoyalFlush(cards []Card) bool {
	// Since a royal flush needs to be straight AND flush, we
	// can utilize the functions defined below for the first part
	if !isFlush(cards) || !isStraight(cards) {
		return false
	}

	// Now we just need to check that the hand includes all royal cards
	ace := false
	king := false
	queen := false
	jack := false
	ten := false

	for _, card := range cards {
		switch card.Rank {
		case 14:
			ace = true
		case 13:
			king = true
		case 12:
			queen = true
		case 11:
			jack = true
		case 10:
			ten = true
		}
	}

	return ace && king && queen && jack && ten
}

// Function to determine if a hand is a flush
func isFlush(cards []Card) bool {
	// Track the number of each suit and determine if
	// there are 5 of any of one of them at the end
	spades := 0
	hearts := 0
	diamonds := 0
	clubs := 0

	for _, card := range cards {
		switch validation.Suits[card.Suit] {
		case "Spades":
			spades++
		case "Hearts":
			hearts++
		case "Diamonds":
			diamonds++
		case "Clubs":
			clubs++
		}
	}

	if spades == 5 || hearts == 5 || diamonds == 5 || clubs == 5 {
		return true
	}

	return false
}

// Function to determine if a hand is a straight
func isStraight(cards []Card) bool {
	for i := 0; i < len(cards)-1; i++ {
		// Since our hand is sorted by highest value first,
		// we can check if the value of each subsequent card
		// is equal to current value minus 1
		if cards[i].Rank != cards[i+1].Rank+1 {
			return false
		}
	}

	return true
}
