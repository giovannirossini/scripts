package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	var name string
	flag.StringVar(&name, "name", "", "Your name")
	var age int
	flag.IntVar(&age, "age", 0, "Your age")

	// Parse flags
	flag.Parse()

	// Use flags in code
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}
