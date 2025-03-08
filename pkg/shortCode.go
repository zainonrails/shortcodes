package shortcodes

import (
	"strings"

	"github.com/zainonrails/shortcodes/internal"
)

// Takes an uint id as an argument which would be taken from a database record
// for which you need to create a shortcode
// example: shortcode.GenerateShortCode(1220) => "QqWZcuh"
func GenerateShortCode(id uint) string {
	var result strings.Builder
	id = internal.ShuffleNumber(id)

	for id > 0 {
		index := id % internal.BASE
		char := internal.CHARACTERS[index]
		result.WriteString(char)

		id /= internal.BASE
	}

	return result.String()
}
