package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	num := 3

	fmt.Print("The number ", num, " is written like this: ")

	switch num {
	case 1:
		fmt.Println("one.")
	case 2:
		fmt.Println("two.")
	case 3:
		fmt.Println("three.")
	}

	fmt.Println("Are you from the Unix family?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Yes!")
	default:
		fmt.Println("Nevermind...")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Ok! Sleep well.")
	default:
		fmt.Println("Wake up!")
	}

	num = 7
	fmt.Println("Does this number fit into one digit?")
	switch {
	case num < 10:
		fmt.Println("Yes!")
	case num >= 10 && num < 100:
		fmt.Println("Does fit into two?")
	case num >= 100:
		fmt.Println("No. This number is too large!")
	}
}
