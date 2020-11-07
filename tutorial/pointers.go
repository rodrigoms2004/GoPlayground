// https://www.golang-book.com/books/intro/3

// go run test.go

// Generate executable
// go build

package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func square(x *float64) {
	*x = *x * *x
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
	fmt.Println(xPtr)  // memory addr from value
	fmt.Println(&xPtr) // memory addr of the pointer

	y := 1.5
	square(&y) // 2.25
	fmt.Println(y)
}
