package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

/*
*
When you have a receiver Go is going to attchatch the method to the type
*/
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()

	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
