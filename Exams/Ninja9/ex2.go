package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

func (p *person) speak() {
	fmt.Println("My name is:", p.First, p.Last)
}

type human interface {
	speak()
}

func saySomething(p human) {
	p.speak()
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   27,
	}

	saySomething(&p1)
}
