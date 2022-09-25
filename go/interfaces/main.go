package main

import (
	"fmt"

	"github.com/giovannirossini/scripts/go/interfaces/models"
)

func main() {
	jojo := models.Bird{}
	jojo.Name = "Jojo"

	chickenSound(jojo)
	duckSound(jojo)
}

func chickenSound(c models.Chicken) {
	fmt.Println(c.Cacarejar())
}

func duckSound(d models.Duck) {
	fmt.Println(d.Quack())
}
