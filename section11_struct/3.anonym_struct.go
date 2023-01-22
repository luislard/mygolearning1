package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1)

	// anonymus struct
	anonym := struct {
		foo string
		bar string
	}{
		foo: "some foo",
		bar: "some bar",
	}

	fmt.Println(anonym)
}
