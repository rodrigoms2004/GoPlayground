package main

import (
	"fmt"
	"time"
)

func feedChannel(c chan int) {
	// aqui vamos feedar nosso canal
	for i := 0; i < 10; i++ {
		c <- i
	}
}

func printFromChannel(c chan int) {
	// nessa função, simplesmente printamos o que está no canal
	for len(c) > 0 {
		fmt.Println(<-c)
	}
}

func main() {
	c := make(chan int, 10) // nosso canal tem tamanho 10

	// feedamos o canal de forma concorrente
	for i := 0; i < 3; i++ {
		go feedChannel(c)
	}

	time.Sleep(5 * time.Second)

	// printamos do canal de forma concorrente
	for i := 0; i < 3; i++ {
		go printFromChannel(c)
	}

	time.Sleep(5 * time.Second)
}
