package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	response, err := client.Get("https://www.google.com")
	if err != nil {
		fmt.Println("[main] There was an error to open Google page. Error: ", err.Error())
		return
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("[main] There was an error to read Google page. Error: ", err.Error())
			return
		}
		fmt.Println("\r\n" + string(body))
	}

	request, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		fmt.Println("[main] There was an error to create a GET request to Google page. Error: ", err.Error())
		return
	}

	request.SetBasicAuth("test", "password")

	response, err = client.Do(request)
	if err != nil {
		fmt.Println("[main] There was an error to open Google page. Error: ", err.Error())
		return
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("[main] There was an error to read Google page. Error: ", err.Error())
			return
		}
		fmt.Println("\r\n" + string(body))
	}
}
