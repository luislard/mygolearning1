package main

import "fmt"

type person struct {
	firstName               string
	lastName                string
	favoriteIcecreamFlavors []string
}

/*
Take the code from the previous exercise, then store the values of type person in a map with the
key of last name. Access each value in the map. Print out the values, ranging over the slice.
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
	//persons := map[string]person{
	//	"Mouse": p1,
	//	"Pluto": p2,
	//}

	persons := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	fmt.Println(persons)

	for key, person := range persons {
		fmt.Println(key)
		personRenderer(person)
	}
}

func personRenderer(person person) {
	fmt.Println(person.firstName)
	fmt.Println(person.lastName)
	for i, flavor := range person.favoriteIcecreamFlavors {
		fmt.Println(i, flavor)
	}
}
