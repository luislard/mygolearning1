package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last)
	fmt.Println("My age is", p.age)
}

func main() {
	p1 := person{
		first: "Luigi",
		last:  "Bros",
		age:   22,
	}

	p1.speak()
}
