package main

import "fmt"

/*
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
TYPE []string which stores their favorite things. Store three records in your map. Print out all of
the values, along with their index position in the slice.
`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
*/
func main() {

	m := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for key, val := range m {
		fmt.Println("These are the favorite things of", key)
		for _, thing := range val {
			fmt.Printf("\t%v\n", thing)
		}
	}

}
