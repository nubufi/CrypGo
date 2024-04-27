package otp

import (
	"math/rand"
)

// RandomSequence generates a random sequence of integers
// with a maximum value of maxInt
// The length of the sequence is equal to the length of the text
// The sequence is used to generate the one-time pad
// for the one-time pad cipher
//
// Parameters:
//
// - text: the text for which the one-time pad is generated
// - maxInt: the maximum integer value for the random sequence
//
// Returns:
//
// - []int: the random sequence
//
// Example:
//
// - RandomSequence("hello", 26)
func RandomSequence(text string, maxInt int) []int {
	var sequence []int
	for range text {
		sequence = append(sequence, rand.Intn(maxInt))
	}
	return sequence
}
