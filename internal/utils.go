package internal

import (
	"math/big"
)

// Simple bijective function
// Shifts all results by a constant
// Makes zero and small numbers look more random
// Makes it harder to guess the original number
// Breaks obvious patterns in the output
// Instead of 'AAAAAAB' or 'AAAAABA'
// produces 'K8mP2nX' and 'R3wQ7vY'

func ShuffleNumber(id *big.Int) *big.Int {
	// keeps it in our range of characters
	// has to be co-prime with primeNumber
	idBig := new(big.Int).Set(id)
	PrimeBig := new(big.Int).Set(PRIME_NUMBER)
	offsetPrimeBig := new(big.Int).Set(OFFSET_PRIME_NUMBER)
	coPrimeNumber := new(big.Int).Exp(BASE, CODE_LENGTH, nil)

	return PrimeBig.Mul(PrimeBig, idBig).Add(PrimeBig, offsetPrimeBig).Mod(PrimeBig, coPrimeNumber)
}

// simple reverse bijective function
// produces original id

func UnShuffleNumber(id *big.Int) uint {
	coPrimeNumber := big.NewInt(0).Exp(BASE, CODE_LENGTH, nil)
	idBig := new(big.Int).Set(id)
	PrimeBig := new(big.Int).Set(PRIME_NUMBER)
	offsetPrimeBig := new(big.Int).Set(OFFSET_PRIME_NUMBER)

	inverse := modInverse(PrimeBig, coPrimeNumber)

	// Perform calculation using big.Int
	result := new(big.Int)
	result.Sub(idBig, offsetPrimeBig)
	result.Add(result, coPrimeNumber)
	result.Mul(result, inverse)
	result.Mod(result, coPrimeNumber)

	return uint(result.Uint64())
}

func modInverse(a, m *big.Int) *big.Int {
	return new(big.Int).ModInverse(a, m)
}
