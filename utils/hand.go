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
func IsStraight(cards []Card) bool {
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
