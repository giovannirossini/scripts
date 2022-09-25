package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/giovannirossini/scripts/go/write_files/models"
)

func main() {
	file, err := os.Open("citys.csv")
	if err != nil {
		fmt.Println("[main] There was an error to open the file. Error: ", err.Error())
		return
	}

	csvReader := csv.NewReader(file)
	content, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("[main] There was an error to read the file. Error: ", err.Error())
		return
	}

	jsonFile, err := os.Create("city.json")
	if err != nil {
		fmt.Println("[main] There was an error to create JSON file. Error: ", err.Error())
		return
	}

	writter := bufio.NewWriter(jsonFile)
	writter.WriteString("[\r\n")

	for _, line := range content {
		for itemIndex, item := range line {
			data := strings.Split(item, "/")
			city := models.City{}
			city.Name = data[0]
			city.State = data[1]
			fmt.Printf("City: %+v\r\n", city)
			cityJSON, err := json.Marshal(city)
			if err != nil {
				fmt.Println("[main] There was an error on item ", item, ". Error:", err.Error())
			}
			writter.WriteString(" " + string(cityJSON))
			if (itemIndex + 1) < len(line) {
				writter.WriteString(",\r\n")
			}
		}
	}

	writter.WriteString("\r\n]")
	writter.Flush()
	jsonFile.Close()
	file.Close()

}
