package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("I = ", i)
	}
	value := 0
	test := true

	for test {
		value++
		if value%5 == 2 {
			test = false
		}
		fmt.Println("Value is: ", value)
	}

	for {
		value--
		fmt.Println("Value is: ", value)
		if value == 0 {
			break
		}
	}

	text := "I love to code using GoLang."

	for index, letter := range text {
		fmt.Printf("Text(%d) = %q\r\n", index, letter)
	}

}
