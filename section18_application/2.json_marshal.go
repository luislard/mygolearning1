package main

import (
	"encoding/json"
	"fmt"
)

// when encoding the properties with first letter in lowercase will be skipped
// when encoding the properties with first letter in uppercase will be encoded
type person struct {
	Name        string
	bankAccount string
}

func main() {
	p1 := person{
		Name:        "James Bond",
		bankAccount: "007",
	}

	p2 := person{
		Name:        "Pippo Pluto",
		bankAccount: "woof",
	}

	people := []person{p1, p2}

	sliceOfBytes, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sliceOfBytes)) // [{"Name":"James Bond"},{"Name":"Pippo Pluto"}]
}
