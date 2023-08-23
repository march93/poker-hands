package main

import (
	"poker-hands/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	input := [][]string{
		{"AH", "JH", "QH", "KH", "TH"},
		{"2H", "3S", "TD", "6C", "JH"},
		{"8H", "8S", "8D", "8C", "TH"},
		{"3H", "6H", "4H", "5H", "7H"},
	}
	hands := evaluate(input)
	assert.Equal(t, hands, []utils.Hand{
		{Cards: []utils.Card{{Rank: 14, Suit: "H"}, {Rank: 13, Suit: "H"}, {Rank: 12, Suit: "H"}, {Rank: 11, Suit: "H"}, {Rank: 10, Suit: "H"}}, Type: 9, Score: 10411194},
		{Cards: []utils.Card{{Rank: 7, Suit: "H"}, {Rank: 6, Suit: "H"}, {Rank: 5, Suit: "H"}, {Rank: 4, Suit: "H"}, {Rank: 3, Suit: "H"}}, Type: 8, Score: 8873283},
		{Cards: []utils.Card{{Rank: 10, Suit: "H"}, {Rank: 8, Suit: "H"}, {Rank: 8, Suit: "S"}, {Rank: 8, Suit: "D"}, {Rank: 8, Suit: "C"}}, Type: 7, Score: 8030344},
		{Cards: []utils.Card{{Rank: 11, Suit: "H"}, {Rank: 10, Suit: "D"}, {Rank: 6, Suit: "C"}, {Rank: 3, Suit: "S"}, {Rank: 2, Suit: "H"}}, Type: 0, Score: 763442},
	})
}
