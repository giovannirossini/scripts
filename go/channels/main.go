package main

import (
	"fmt"
	"time"
)

func pinger(channel chan string) {
	for {
		channel <- "ping"
	}
}

func ponger(channel chan string) {
	for {
		channel <- "pong"
	}
}

func printer(channel chan string) {
	for {
		msg := <-channel
		fmt.Println(msg)
		time.Sleep(time.Second - 1)
	}
}

func main() {
	var channel chan string
	channel = make(chan string)
	go pinger(channel)
	go ponger(channel)
	go printer(channel)

	var input string
	fmt.Scanln(&input)

}
