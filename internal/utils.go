package internal

import (
	"math"
)

// Simple bijective function
// You can adjust these prime numbers for different patterns
// Shifts all results by a constant
// Makes zero and small numbers look more random
// Makes it harder to guess the original number
// Breaks obvious patterns in the output
// Instead of 'AAAAAAB' or 'AAAAABA'
// produces 'K8mP2nX' and 'R3wQ7vY'
func ShuffleNumber(id uint) uint {
	// has to be prime
	var a uint = 33_33_33_33_33
	// doesn't have to be prime
	var b uint = 77_77_77_77_77
	// keeps it in our range of characters
	var m uint = uint(math.Pow(BASE, CODE_LENGTH))

	return ((a * id) + b) % m
}
