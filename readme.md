# Shortcodes

A lightweight Go package that generates unique shortcodes from unsigned integer IDs.

## Installation

```bash
go get github.com/zainonrails/shortcodes
```

## Usage

```go
package main

import (
    "fmt"
	shortcodes "github.com/zainonrails/shortcodes/pkg"
)

func main() {
    // Generate a shortcode from a uint ID
	shortcode := shortcodes.GenerateShortCode(1233445)
	fmt.Println(shortcode)

	original := shortcodes.DecodeShortCode(shortcode)
	fmt.Println(original)
    
    // Use in your database operations
    // dbRecord.ShortCode = shortcode.GenerateShortCode(dbRecord.ID)
}
```

## Features

- Generates URL-safe shortcodes
- Deterministic output (same ID always generates same shortcode)
- Decodes back to original ID
- No external dependencies
- Thread-safe

## How It Works

The package uses a bijective function to create shortcodes:

- Applies carefully selected prime numbers to transform the original integer
- Shifts all results by a constant value to make even small numbers look random
- Makes it harder to reverse-engineer the original ID
- Breaks obvious sequential patterns in the output

For example, instead of sequential IDs generating predictable patterns like 'AAAAAAB' or 'AAAAABA',
the algorithm produces randomized-looking codes like 'K8mP2nX' and 'R3wQ7vY'.
This approach provides a good balance between uniqueness, readability, and security through obscurity
while maintaining a consistent length for IDs in the same numeric range.

## Use cases

- **URL Shortening**: Create compact, user-friendly URLs from database record IDs
- **Referral Codes**: Generate memorable invitation or referral codes
- **Product SKUs**: Create readable, unique identifiers for products
- **Public API IDs**: Obscure sequential database IDs in public-facing APIs
- **Coupon Codes**: Generate unique, typeable discount codes
- **Ticket/Order Numbers**: Create human-readable identifiers for orders or support tickets
- **Obfuscation**: Hide the actual count or sequence of your database records from users

## API

The package provides two methods:
```go
GenerateShortCode(id uint) string  // generates 7 digit code by default
```
Converts an unsigned integer ID into a unique alphanumeric shortcode

```go
DecodeShortCode(code string) uint
```
Decodes string code back to original ID
