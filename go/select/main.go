package main

import (
	"fmt"
	"time"
)

var irc = make(chan string)
var sms = make(chan string)

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

func whazup(channel chan string) {
	for {
		channel <- "Whazup??"
	}
}

func heyThere(channel chan string) {
	for {
		channel <- "Hey there!"
	}
}

func printer() {
	var msg string
	for {
		select {
		case msg = <-irc:
			fmt.Println(msg)
			time.Sleep(time.Second - 1)
		case msg = <-sms:
			fmt.Println("ZzZ....ZzzZ ", msg)
		}
		time.Sleep(time.Second - 1)
	}
}

func main() {
	go pinger(irc)
	go ponger(irc)
	go whazup(sms)
	go heyThere(sms)
	go printer()

	var input string
	fmt.Scanln(&input)

}
