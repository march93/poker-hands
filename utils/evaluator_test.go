package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTransform(t *testing.T) {
	cards := sortCards([]Card{
		{Rank: 2, Suit: "S"},
		{Rank: 10, Suit: "H"},
		{Rank: 3, Suit: "D"},
		{Rank: 14, Suit: "C"},
		{Rank: 8, Suit: "H"},
	})
	value := BinaryTransform(cards)
	assert.Equal(t, value, 960562)

	cards = sortCards([]Card{
		{Rank: 12, Suit: "S"},
		{Rank: 10, Suit: "S"},
		{Rank: 11, Suit: "S"},
		{Rank: 14, Suit: "S"},
		{Rank: 13, Suit: "S"},
	})
	value = BinaryTransform(cards)
	assert.Equal(t, value, 10411194)
}
