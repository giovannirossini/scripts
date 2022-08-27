package basic

import "fmt"

// Realty is a struct that store data of a realty
type Realty struct {
	X     int
	Y     int
	Name  string
	value int
}

func main() {
	house := Realty{}
	fmt.Printf("The house is: %+v\r\n", house)

	apartment := Realty{17, 56, "My apartment", 760000}
	fmt.Printf("The apartment is: %+v\r\n", apartment)

	farm := Realty{
		Y:     85,
		Name:  "My farm",
		value: 1200000,
		X:     144,
	}

	fmt.Printf("The farm is: %+v\r\n", farm)

	house.Name = "Home Sweet Home"
	house.value = 450000
	house.X = 8
	house.Y = 7
	fmt.Printf("The house is: %+v\r\n", house)
}
