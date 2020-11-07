// https://www.golang-book.com/books/intro/3

// go run test.go

// Generate executable
// go build

package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}
func (p *Person) setName(name string) {
	p.Name = name
}

type Android struct {
	Person // same as 'Person Person'
	Model  string
}

func main() {
	a := new(Android)
	a.setName("Rodrigo")
	a.Talk()
	a.Person.Talk()
}
