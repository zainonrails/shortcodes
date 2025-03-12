package main

import (
	"fmt"

	shortcodes "github.com/zainonrails/shortcodes/pkg"
)

func main() {
	shortcode := shortcodes.GenerateShortCode(1220)
	fmt.Println(shortcode)
}
