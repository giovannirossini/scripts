package main

import (
	"encoding/json"
	"fmt"

	"github.com/giovannirossini/curso/basic/structs/advanced/models"
)

func main() {
	house := models.Realty{}
	house.Name = "Yellow House"
	house.X = 18
	house.Y = 25
	house.SetValue(60000)

	fmt.Printf("House is: %+v\r\n", house)
	fmt.Printf("The House value is: %+v\r\n", house.GetValue())

	objJSON, _ := json.Marshal(house)
	fmt.Println("House in JSON: ", string(objJSON))
}
