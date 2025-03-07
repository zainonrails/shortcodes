package shortcodes

import (
	"math"
	"strings"
)

const BASE = 62
const CODE_LENGTH = 7

var CHARACTERS = [62]string{"p", "X", "4", "k", "V", "b", "N", "8", "i", "Q", "w", "Z", "2", "F", "m", "Y", "j", "U", "9", "R", "d", "H", "5", "L", "n", "C", "a", "W", "1", "G", "y", "D", "e", "0", "S", "h", "B", "7", "t", "O", "c", "I", "v", "3", "M", "q", "T", "r", "6", "E", "K", "f", "A", "x", "J", "g", "P", "s", "l", "u", "o", "z"}

func GenerateShortCode(id uint) string {
	var result strings.Builder
	id = shuffleNumber(id)

	for id > 0 {
		index := id % BASE
		char := CHARACTERS[index]
		result.WriteString(char)

		id /= BASE
	}

	return result.String()
}

func shuffleNumber(id uint) uint {
	var a uint = 33_33_33_33_33
	var b uint = 77_77_77_77_77
	var m uint = uint(math.Pow(BASE, CODE_LENGTH))

	return ((a * id) + b) % m
}
