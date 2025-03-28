package main

import (
	"fmt"
	"math/big"

	shortcodes "github.com/zainonrails/shortcodes/pkg"
)

func main() {
	shortcode := shortcodes.GenerateShortCode(big.NewInt(1233445))
	fmt.Println(shortcode)

	original := shortcodes.DecodeShortCode(shortcode)
	fmt.Println(original)
}
