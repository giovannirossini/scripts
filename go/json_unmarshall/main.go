package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/giovannirossini/scripts/go/json_unmarshall/models"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	response, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("[main] There was an error to GET JSON Placeholder. Error: ", err.Error())
		return
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("[main] There was an error to read JSON Placeholder. Error: ", err.Error())
			return
		}

		post := models.BlogPost{}
		err = json.Unmarshal(body, &post)
		if err != nil {
			fmt.Println("[main] There was an error to convert the JSON from server. Error: ", err.Error())
			return
		}
		fmt.Printf("Post is: %+v\r\n", post)
	}
}
