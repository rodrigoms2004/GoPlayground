// https://www.golang-book.com/books/intro/3

// go run test.go

// Generate executable
// go build

package main

import (
	"fmt"
	"math"
)

// Structs
type Circle struct {
	x, y, r float64
}

// Fields
func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func circleAreaReference(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

// Methods
// use c.area()
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {
	// var c Circle // {0 0 0}
	// c := new(Circle) // *c is {0 0 0}
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(c))
	fmt.Println(circleAreaReference(&c))
	fmt.Println(c.area())
	fmt.Println("Rectangle")
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}
