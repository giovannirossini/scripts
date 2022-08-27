package main

import "fmt"

func main() {
	month := 0
	status := true
	city := "Chicago"

	if month <= 6 {
		fmt.Println("This debt is recent.")
	}

	if status {
		fmt.Println("Debtor")
	}

	if city == "Chicago" {
		fmt.Println(city)
	}

	if description, status := debtTime(month); status {
		fmt.Println("Situation:", description)
		return
	}

	fmt.Println("Thank you for consulting us!")
}

// DebtTime to describe if a client is a debtor
func debtTime(month int) (description string, status bool) {
	if month > 0 {
		status = true
		description = "The client is a debtor."
		return
	}
	description = "The client has no debt."
	return
}
