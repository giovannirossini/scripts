package main

import (
	"fmt"

	"github.com/giovannirossini/scripts/go/structs/basic"
)

func main() {
	house := new(basic.Realty)
	fmt.Printf("House is: %p - %+v\r\n", &house, house)

	farm := basic.Realty{
		Name: "My Farm",
		X:    12,
		Y:    18,
	}
	fmt.Printf("The farm is: %p - %+v\r\n", &farm, farm)

	changeRealty(&farm)
	fmt.Printf("The new farm propertys is: %p - %+v\r\n", &farm, farm)
}

func changeRealty(realty *basic.Realty) {
	realty.X = realty.X + 10
	realty.Y = realty.Y - 5
}
