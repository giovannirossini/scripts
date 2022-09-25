package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/giovannirossini/scripts/go/web_post/models"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	userRequest := models.User{}
	userRequest.ID = 1
	userRequest.Name = "Wade"

	content, err := json.Marshal(userRequest)
	if err != nil {
		fmt.Println("[main] There was an error on JSON of user object. Error: ", err.Error())
		return
	}

	request, err := http.NewRequest("POST", "http://requestbin.fullcontact.com/qknmb1qk", bytes.NewBuffer(content))
	if err != nil {
		fmt.Println("[main] There was an error to create a POST request to Request Bin. Error: ", err.Error())
		return
	}

	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/json; charset=utf-8")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[main] There was an error to call RequestBin. Error: ", err.Error())
		return
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("[main] There was an error to return RequestBin content type. Error: ", err.Error())
			return
		}
		fmt.Println(string(body))
	}
}
