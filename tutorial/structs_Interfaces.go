// https://www.golang-book.com/books/intro/9

// go run test.go

// Generate executable
// go build

package main

import (
	"fmt"
	"math"
)

// Circle
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return math.Pi * c.r * 2
}

// Rectangle
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

func (r *Rectangle) perimeter() float64 {
	return r.x1 + r.y1 + r.x2 + r.y2
}

// Interface
type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

// Multishape
type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) perimeter() float64 {
	var perimeter float64
	for _, s := range m.shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}

	// m := MultiShape{}
	// m.shapes = []Shape{&c, &r}
	m := MultiShape{[]Shape{&c, &r}}

	fmt.Println("\nAREA:")
	fmt.Println("Circle Area", c.area())
	fmt.Println("Rectangle Area", r.area())
	fmt.Println("Total Area using interface", totalArea(&c, &r))
	fmt.Println("Total Area using MultiShape", m.area())

	fmt.Println("\nPERIMETER:")
	fmt.Println("Circle perimeter", c.perimeter())
	fmt.Println("Rectangle perimeter", r.perimeter())
	fmt.Println("Total Perimeter using interface", totalPerimeter(&c, &r))

	fmt.Println("Total Perimeter using MultiShape", m.perimeter())
}
