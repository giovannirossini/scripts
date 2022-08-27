package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("citys.csv")
	if err != nil {
		fmt.Println("[main] There was an error to open the file. Error: ", err.Error())
		return
	}

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Println("The file content is: ", line)
	// }

	csvReader := csv.NewReader(file)
	content, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("[main] There was an error to read the file. Error: ", err.Error())
		return
	}

	for _, line := range content {
		fmt.Printf("Line is %s\r\n", line)
		for itemIndex, item := range line {
			fmt.Printf("Item [%d] is: %s\r\n", itemIndex, item)
		}
	}
	file.Close()

}
