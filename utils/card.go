package utils

import (
	"poker-hands/validation"
	"sort"
)

type Card struct {
	Rank int
	Suit string
}

// Transform each string representation
// of a card into a Card struct
func Transform(input []string) []Card {
	cards := []Card{}

	for _, card := range input {
		// Get the value of the card as an int
		// mapped to the Ranks map in validator
		rank := validation.Ranks[card[0:1]]
		suit := card[1:]

		s := Card{rank, suit}
		cards = append(cards, s)
	}

	return sortCards(cards)
}

// Sort each hand in descending order by highest ranked card
func sortCards(hand []Card) []Card {
	sort.SliceStable(hand, func(i, j int) bool {
		return hand[i].Rank > hand[j].Rank
	})

	return hand
}
