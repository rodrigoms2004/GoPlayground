// https://www.golang-book.com/books/intro/10

// go run test.go

// Generate executable
// go build

package main

import (
	"fmt"
	"time"
)

// send to receive-only type <-chan string
// receive from send-only type chan<- string)

// Receiver
// func pinger(c chan<- string) { or bi-directional:
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// Sender
// func printer(c <-chan string) { or bi-directional:
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
