package main

import (
	"fmt"
	shortcodes "github.com/zainonrails/shortcodes/pkg"
)

func main() {
	shortcode := shortcodes.GenerateShortCode(1233445)
	fmt.Println(shortcode)

	original := shortcodes.DecodeShortCode(shortcode)
	fmt.Println(original)
}
