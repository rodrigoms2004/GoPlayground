// https://www.golang-book.com/books/intro/3

// go run test.go

// Generate executable
// go build

package main

import "fmt"

func main() {
	xs := []float64{98, 93, 77, 82, 83}

	fmt.Println(average(xs))
	fmt.Println(add(1, 2, 3))

	//  Closure
	// It is possible to create functions inside of functions:
	add2 := func(x, y int) int {
		return x + y
	}
	fmt.Println(add2(1, 1))

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

	fmt.Println("Factorial of 5: ", factorial(5))
	fmt.Println("\n")

	defer second()
	first()
	fmt.Println("\n")

	defer func() {
		str := recover()
		fmt.Println(str, "XPTO")
	}()
	panic("PANIC")

}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// Recursion
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func average(xs []float64) float64 {

	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// Variadic Functions
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
