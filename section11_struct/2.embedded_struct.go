package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool // license to kill
}

func main() {
	secretAgent := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   28,
	}
	fmt.Println(secretAgent)
	fmt.Println(p2)
	// notice that the props got promoted to the outer struct
	fmt.Println(secretAgent.first, secretAgent.last, secretAgent.age, secretAgent.ltk)
	// you can still refer to the inner prop
	fmt.Println(secretAgent.person.first, secretAgent.person.last, secretAgent.person.age, secretAgent.ltk)
	fmt.Println(p2.first, p2.last)
}
