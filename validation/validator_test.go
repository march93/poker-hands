package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	// Empty list
	cards := []string{}
	err := Validate(cards)
	assert.Equal(t, err.Error(), "require 5 cards, 0 card(s) given")

	// Not enough cards
	cards = []string{"2H"}
	err = Validate(cards)
	assert.Equal(t, err.Error(), "require 5 cards, 1 card(s) given")

	// Invalid card length
	cards = []string{"123", "2H", "5C", "6D", "QC"}
	err = Validate(cards)
	assert.Equal(t, err.Error(), "invalid card format provided: 123")

	// Invalid rank
	cards = []string{"KH", "9S", "4C", "0D", "TD"}
	err = Validate(cards)
	assert.Equal(t, err.Error(), "0D - invalid rank provided: 0")

	// Invalid suit
	cards = []string{"KH", "9S", "4C", "JP", "TD"}
	err = Validate(cards)
	assert.Equal(t, err.Error(), "JP - invalid suit provided: P")

	// Valid cards
	cards = []string{"KH", "9S", "4C", "QD", "TD"}
	err = Validate(cards)
	assert.Equal(t, err, nil)
}
