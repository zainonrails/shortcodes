package shortcodes

import (
	"math/big"
	"strings"

	"github.com/zainonrails/shortcodes/internal"
)

// Takes an uint as an argument which would be taken from a database record
// for which you need to create a shortcode
// example: shortcode.GenerateShortCode(big.NewInt(1220)) => "QqWZcuh"

func GenerateShortCode(id uint) string {
	idBig := big.NewInt(int64(id))
	var result strings.Builder
	idBig = internal.ShuffleNumber(idBig)

	// converts to base 62
	for idBig.Cmp(big.NewInt(0)) == 1 {
		index := big.NewInt(0).Mod(idBig, internal.BASE).Int64()
		char := internal.CHARACTERS[index]
		result.WriteString(char)
		idBig.Div(idBig, internal.BASE)
	}

	// reverse string to make significant bit on the left
	var reverseString strings.Builder
	for i := len(result.String()) - 1; i >= 0; i-- {
		reverseString.WriteString(string(result.String()[i]))
	}

	return reverseString.String()
}

func DecodeShortCode(code string) uint {
	var characters = strings.Join(internal.CHARACTERS, "")
	id := big.NewInt(0)

	// converts to base 10
	for _, char := range code {
		idx := strings.Index(characters, string(char))
		id.Mul(id, internal.BASE)
		id.Add(id, big.NewInt(int64(idx)))
	}

	return internal.UnShuffleNumber(id)
}
