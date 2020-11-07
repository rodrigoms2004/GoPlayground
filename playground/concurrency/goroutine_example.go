package main

// go run goroutine_example.go

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello world goroutine")
}

//  função hello roda de forma concorrente e independente da main
func main() {
	// para rodar uma função como goroutine, basta escrever go antes de chama-la
	go hello()
	fmt.Println("main function")

	var input string
	fmt.Scanln(&input)
}
