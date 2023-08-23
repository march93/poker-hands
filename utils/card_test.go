package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortCards(t *testing.T) {
	cards := []Card{
		{Rank: 2, Suit: "S"},
		{Rank: 10, Suit: "H"},
		{Rank: 3, Suit: "D"},
		{Rank: 14, Suit: "C"},
		{Rank: 8, Suit: "H"},
	}
	sorted := sortCards(cards)
	assert.Equal(t, sorted, []Card{
		{Rank: 14, Suit: "C"},
		{Rank: 10, Suit: "H"},
		{Rank: 8, Suit: "H"},
		{Rank: 3, Suit: "D"},
		{Rank: 2, Suit: "S"},
	})
}

func TestTransform(t *testing.T) {
	cards := []string{"2S", "TH", "3D", "AC", "8H"}
	sorted := Transform(cards)
	assert.Equal(t, sorted, []Card{
		{Rank: 14, Suit: "C"},
		{Rank: 10, Suit: "H"},
		{Rank: 8, Suit: "H"},
		{Rank: 3, Suit: "D"},
		{Rank: 2, Suit: "S"},
	})
}
