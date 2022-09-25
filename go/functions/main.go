package main

import (
	"fmt"

	"github.com/giovannirossini/scripts/go/functions/math"
)

func main() {
	result := math.Calculation(math.Divide, 2, 2)
	fmt.Printf("Divide result is: %d\r\n", result)
	result = math.Calculation(math.Minus, 16, 3)
	fmt.Printf("Minus result is: %d\r\n", result)
	result, rest := math.DivideWithRest(16, 3)
	fmt.Printf("Divide with rest result is: %d and rest: %d\r\n", result, rest)
}
