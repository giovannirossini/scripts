package main

import "fmt"

var (
	// Address var is public and var phone is private
	// Capitalize vars are public
	Address, phone    string
	amount, inventory int
	buy               bool
	price             float64
	word              rune
)

func main() {
	fmt.Printf("Address: %s\r\n", Address)
	fmt.Printf("Phone: %s\r\n", phone)
	fmt.Printf("Amount: %d\r\n", amount)
	fmt.Printf("Inventory: %d\r\n", inventory)
	fmt.Printf("Buy: %v\r\n", buy)
	fmt.Printf("Price: %f\r\n", price)
	fmt.Printf("Word: %v\r\n", word)
}
