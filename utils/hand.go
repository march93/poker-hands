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

// Function to determine if a hand is a straight flush
func isStraightFlush(cards []Card) bool {
	// Check for flush and straight
	return isFlush(cards) && isStraight(cards)
}

// Function to determine if a hand is a four of a kind.
// Instead of double looping through every combination,
// we consider a different approach
func isFourOfAKind(cards []Card) bool {
	// Since our cards are already sorted, there are only
	// two ways we can have a four of a kind setup
	// [A A A A 3] or [A 3 3 3 3] front or back
	// So we just need to check these two conditions
	front := (cards[0].Rank == cards[1].Rank) && (cards[0].Rank == cards[2].Rank) && (cards[0].Rank == cards[3].Rank)
	back := (cards[4].Rank == cards[3].Rank) && (cards[4].Rank == cards[2].Rank) && (cards[4].Rank == cards[1].Rank)

	return front || back
}

// Function to determine if a hand is a full house.
// We also consider a different approach vs double loops
func isFullHouse(cards []Card) bool {
	// Since cards are sorted, there are only two ways
	// to form a full house with our hand.
	// [X X X Y Y] or [X X Y Y Y]
	front := (cards[0].Rank == cards[1].Rank) && (cards[0].Rank == cards[2].Rank) && (cards[0].Rank != cards[3].Rank) && (cards[3].Rank == cards[4].Rank)
	back := (cards[0].Rank == cards[1].Rank) && (cards[0].Rank != cards[2].Rank) && (cards[2].Rank == cards[3].Rank) && (cards[2].Rank == cards[4].Rank)

	return front || back
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

// Function to determine if a hand is a three of a kind
func isThreeOfAKind(cards []Card) bool {
	// Use a map to track rank occurrences and
	// the largest total occurrence so far
	rankMap := make(map[int]int)
	max := 0

	for _, card := range cards {
		// Fetch rank total from the map
		total, ok := rankMap[card.Rank]

		if ok {
			// If it exists in the map, increment total
			rankMap[card.Rank] = total + 1
		} else {
			// Otherwise, create new entry
			rankMap[card.Rank] = 1
		}

		// Update max if new value is larger
		if rankMap[card.Rank] > max {
			max = rankMap[card.Rank]
		}
	}

	return max == 3
}

// Function to determine if a hand is a two pair
func isTwoPair(cards []Card) bool {
	// Use a map to track rank occurrences and
	// the largest occurrence so far
	rankMap := make(map[int]int)
	max := 0

	for _, card := range cards {
		// Fetch rank total from the map
		total, ok := rankMap[card.Rank]

		if ok {
			// If it exists in the map, increment total
			rankMap[card.Rank] = total + 1
		} else {
			// Otherwise, create new entry
			rankMap[card.Rank] = 1
		}

		// Update max if new value is larger
		if rankMap[card.Rank] > max {
			max = rankMap[card.Rank]
		}
	}

	// For two pair, we must ensure that there are 3
	// total unique ranks and the largest occurrence is 2
	return len(rankMap) == 3 && max == 2
}

// Function to determine if a hand is a pair
func isPair(cards []Card) bool {
	// Use a map to track rank occurrences and
	// the largest occurrence so far
	rankMap := make(map[int]int)
	max := 0

	for _, card := range cards {
		// Fetch rank total from the map
		total, ok := rankMap[card.Rank]

		if ok {
			// If it exists in the map, increment total
			rankMap[card.Rank] = total + 1
		} else {
			// Otherwise, create new entry
			rankMap[card.Rank] = 1
		}

		// Update max if new value is larger
		if rankMap[card.Rank] > max {
			max = rankMap[card.Rank]
		}
	}

	// For a pair, we must ensure that there are 4
	// total unique ranks and the largest occurrence is 2
	return len(rankMap) == 4 && max == 2
}
