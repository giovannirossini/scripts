package main

import (
	"fmt"

	"github.com/giovannirossini/curso/basic/maps/models"
)

var cache map[string]models.Realty

func main() {

	cache = make(map[string]models.Realty, 0)

	house := models.Realty{}
	house.Name = "Yellow House"
	house.X = 18
	house.Y = 25
	house.SetValue(60000)

	cache["Yellow House"] = house

	fmt.Println("It has a yellow house in the cache?")
	realty, find := cache["Yellow House"]
	if find {
		fmt.Printf("Yes: %+v\r\n", realty)
	}

	apartment := models.Realty{}
	apartment.Name = "Green Apartment"
	apartment.X = 23
	apartment.Y = 55
	apartment.SetValue(160000)

	cache[apartment.Name] = apartment

	fmt.Println("How many itens has the cache?", len(cache))

	for key, realty := range cache {
		fmt.Printf("Key[%s] = %+v\r\n", key, realty)
	}

	_, find = cache[house.Name]
	if find {
		delete(cache, realty.Name)
	}

	fmt.Println("How many itens has the cache now?", len(cache))
}
