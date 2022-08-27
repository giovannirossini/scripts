package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/giovannirossini/curso/intermediary/goroutines/models"
)

var orchestrator sync.WaitGroup

// Test with two go and one go, and no go command.
// With two goroutines we have the concurrent jobs
// CONCURRENT is different from PARALLEL
func main() {
	orchestrator.Add(2)
	go translateJSON("citys")
	go translateJSON("citys2")
	orchestrator.Wait()
}

func translateJSON(fileName string) {
	fmt.Println(time.Now(), " - Start to translate file: ", fileName)
	file, err := os.Open(fileName + ".csv")
	if err != nil {
		fmt.Println("[main] There was an error to open the file. Error: ", err.Error())
		return
	}

	// Defer to handle with mem and connection
	defer file.Close()

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

	defer jsonFile.Close()

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
	fmt.Println(time.Now(), " - The translations of the file: ", fileName, " was finished.")
	orchestrator.Done()
}
