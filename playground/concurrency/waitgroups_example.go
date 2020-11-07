package main

import (
	"fmt"
	"sync"
)

func feedChannel(c chan int, wg *sync.WaitGroup) {
	// aqui vamos feedar nosso canal
	for i := 0; i < 10; i++ {
		c <- i
	}
	wg.Done()
}

func printFromChannel(c chan int, wg *sync.WaitGroup) {
	// nessa função, simplesmente printamos o que está no canal
	for len(c) > 0 {
		fmt.Println(<-c)
	}
	wg.Done()
}

func main() {
	c := make(chan int, 100) // nosso canal tem tamanho 100
	var wg sync.WaitGroup    // nosso wait_group

	// feedamos o canal de forma concorrente
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go feedChannel(c, &wg)
	}
	wg.Wait()

	// printamos do canal de forma concorrente
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go printFromChannel(c, &wg)
	}
	wg.Wait()

}
