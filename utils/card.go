package utils

type Card struct {
	Rank string
	Suit string
}

func Transform(input []string) []Card {
	cards := []Card{}

	for _, card := range input {
		rank := card[0:1]
		suit := card[1:]

		s := Card{rank, suit}
		cards = append(cards, s)
	}

	return cards
}
