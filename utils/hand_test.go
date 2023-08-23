package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Pre-initialized hands
var high = []Card{
	{Rank: 2, Suit: "S"},
	{Rank: 10, Suit: "H"},
	{Rank: 3, Suit: "D"},
	{Rank: 14, Suit: "C"},
	{Rank: 8, Suit: "H"},
}

var pair = []Card{
	{Rank: 2, Suit: "S"},
	{Rank: 10, Suit: "H"},
	{Rank: 3, Suit: "D"},
	{Rank: 14, Suit: "C"},
	{Rank: 2, Suit: "H"},
}

var twoPair = []Card{
	{Rank: 2, Suit: "S"},
	{Rank: 10, Suit: "H"},
	{Rank: 3, Suit: "D"},
	{Rank: 10, Suit: "C"},
	{Rank: 2, Suit: "H"},
}

var threeOfAKind = []Card{
	{Rank: 2, Suit: "S"},
	{Rank: 10, Suit: "H"},
	{Rank: 2, Suit: "D"},
	{Rank: 14, Suit: "C"},
	{Rank: 2, Suit: "H"},
}

var straight = []Card{
	{Rank: 2, Suit: "S"},
	{Rank: 5, Suit: "H"},
	{Rank: 3, Suit: "D"},
	{Rank: 4, Suit: "C"},
	{Rank: 6, Suit: "H"},
}

var flush = []Card{
	{Rank: 2, Suit: "S"},
	{Rank: 10, Suit: "S"},
	{Rank: 3, Suit: "S"},
	{Rank: 14, Suit: "S"},
	{Rank: 9, Suit: "S"},
}

var fullHouse = []Card{
	{Rank: 2, Suit: "S"},
	{Rank: 10, Suit: "H"},
	{Rank: 2, Suit: "D"},
	{Rank: 10, Suit: "C"},
	{Rank: 2, Suit: "H"},
}

var fourOfAKind = []Card{
	{Rank: 2, Suit: "S"},
	{Rank: 2, Suit: "H"},
	{Rank: 2, Suit: "D"},
	{Rank: 14, Suit: "C"},
	{Rank: 2, Suit: "C"},
}

var straightFlush = []Card{
	{Rank: 2, Suit: "S"},
	{Rank: 5, Suit: "S"},
	{Rank: 3, Suit: "S"},
	{Rank: 4, Suit: "S"},
	{Rank: 6, Suit: "S"},
}

var royalFlush = []Card{
	{Rank: 10, Suit: "S"},
	{Rank: 14, Suit: "S"},
	{Rank: 12, Suit: "S"},
	{Rank: 13, Suit: "S"},
	{Rank: 11, Suit: "S"},
}

func TestIsPair(t *testing.T) {
	valid := isPair(sortCards(pair))
	assert.Equal(t, valid, true)

	valid = isPair(sortCards(twoPair))
	assert.Equal(t, valid, false)
}

func TestIsTwoPair(t *testing.T) {
	valid := isTwoPair(sortCards(twoPair))
	assert.Equal(t, valid, true)

	valid = isTwoPair(sortCards(pair))
	assert.Equal(t, valid, false)
}

func TestIsThreeOfAKind(t *testing.T) {
	valid := isThreeOfAKind(sortCards(threeOfAKind))
	assert.Equal(t, valid, true)

	valid = isThreeOfAKind(sortCards(fourOfAKind))
	assert.Equal(t, valid, false)
}

func TestIsStraight(t *testing.T) {
	valid := isStraight(sortCards(straight))
	assert.Equal(t, valid, true)

	valid = isStraight(sortCards(high))
	assert.Equal(t, valid, false)
}

func TestIsFlush(t *testing.T) {
	valid := isFlush(sortCards(flush))
	assert.Equal(t, valid, true)

	valid = isFlush(sortCards(straight))
	assert.Equal(t, valid, false)
}

func TestIsFullHouse(t *testing.T) {
	valid := isFullHouse(sortCards(fullHouse))
	assert.Equal(t, valid, true)

	valid = isFullHouse(sortCards(threeOfAKind))
	assert.Equal(t, valid, false)
}

func TestIsFourOfAKind(t *testing.T) {
	valid := isFourOfAKind(sortCards(fourOfAKind))
	assert.Equal(t, valid, true)

	valid = isFourOfAKind(sortCards(threeOfAKind))
	assert.Equal(t, valid, false)
}

func TestIsStraightFlush(t *testing.T) {
	valid := isStraightFlush(sortCards(straightFlush))
	assert.Equal(t, valid, true)

	valid = isStraightFlush(sortCards(straight))
	assert.Equal(t, valid, false)
}

func TestIsRoyalFlush(t *testing.T) {
	valid := isRoyalFlush(sortCards(royalFlush))
	assert.Equal(t, valid, true)

	valid = isRoyalFlush(sortCards(fourOfAKind))
	assert.Equal(t, valid, false)
}

func TestDetermineHand(t *testing.T) {
	hand := DetermineHand(sortCards(high))
	assert.Equal(t, hand, 0)

	hand = DetermineHand(sortCards(pair))
	assert.Equal(t, hand, 1)

	hand = DetermineHand(sortCards(twoPair))
	assert.Equal(t, hand, 2)

	hand = DetermineHand(sortCards(threeOfAKind))
	assert.Equal(t, hand, 3)

	hand = DetermineHand(sortCards(straight))
	assert.Equal(t, hand, 4)

	hand = DetermineHand(sortCards(flush))
	assert.Equal(t, hand, 5)

	hand = DetermineHand(sortCards(fullHouse))
	assert.Equal(t, hand, 6)

	hand = DetermineHand(sortCards(fourOfAKind))
	assert.Equal(t, hand, 7)

	hand = DetermineHand(sortCards(straightFlush))
	assert.Equal(t, hand, 8)

	hand = DetermineHand(sortCards(royalFlush))
	assert.Equal(t, hand, 9)
}
