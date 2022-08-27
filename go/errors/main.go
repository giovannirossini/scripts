package main

import (
	"encoding/json"
	"fmt"

	"github.com/giovannirossini/curso/basic/errors/models"
)

func main() {
	house := models.Realty{}
	house.Name = "Yellow House"
	house.X = 18
	house.Y = 25
	if err := house.SetValue(40000000); err != nil {
		fmt.Println("[main] There was an error on realty value: ", err.Error())
		if err == models.ErrorValueRealty {
			fmt.Println("Broker check your evaluation: ", err.Error())
		}
		return
	}

	fmt.Printf("House is: %+v\r\n", house)
	fmt.Printf("The House value is: %+v\r\n", house.GetValue())

	objJSON, err := json.Marshal(house)
	if err != nil {
		fmt.Println("[main] There was an error generating JSON.", err.Error())
		return
	}
	fmt.Println("House in JSON: ", string(objJSON))
}
