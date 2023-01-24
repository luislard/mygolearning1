package main

import "fmt"

type person struct {
	firstName               string
	lastName                string
	favoriteIcecreamFlavors []string
}

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:
● first name
● last name
● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
which stores the favorite flavors
*/
func main() {
	p1 := person{
		firstName: "Mickey",
		lastName:  "Mouse",
		favoriteIcecreamFlavors: []string{
			"chocolat",
			"haznutzl",
		},
	}

	p2 := person{
		firstName: "Pippo",
		lastName:  "Pluto",
		favoriteIcecreamFlavors: []string{
			"haznutzl",
		},
	}
	persons := []person{p1, p2}
	fmt.Println(persons)

	for _, person := range persons {
		fmt.Println(person.firstName)
		fmt.Println(person.lastName)
		for i, flavor := range person.favoriteIcecreamFlavors {
			fmt.Println(i, flavor)
		}
	}
}
