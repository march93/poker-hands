package utils

/* Transform hand via a binary representation, where
 * the hand rank comes first, then the card values.
 * e.g. 1001 1110 1101 1100 1011 1010 -> 10411194
 */
func BinaryTransform(hand []Card) int {
	// value := integerToBinary(determineHand(hand))
	value := DetermineHand(hand)

	for _, card := range hand {
		// Shift current value 4 bits to the left so
		// we can add the rest of the card bits.
		// This should give us a unique representation
		// of the hand value since it's already sorted
		value = (value << 4) + card.Rank
	}

	return value
}
