package main

import "fmt"

func main() {

	var array [3]int
	array[0] = 3
	array[1] = 2
	array[2] = 1

	fmt.Println("What is the capacity of this array?", len(array))

	singers := [2]string{"Ice Blue", "Drake"}
	fmt.Printf("What is into this array? \n\r%v\r\n", singers)

	citys := [...]string{"Chicago", "Los Angeles", "São Paulo", "Luanda", "Maputo"}
	fmt.Printf("What is into this array? \n\r%v\r\n", citys)
	for _, city := range citys {
		fmt.Printf("This town is called: %s\r\n", city)
	}
}
