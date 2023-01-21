package main

import "fmt"

/*
delete a record from the map
*/
func main() {

	m := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	delete(m, "no_dr")

	for key, val := range m {
		fmt.Println("These are the favorite things of", key)
		for _, thing := range val {
			fmt.Printf("\t%v\n", thing)
		}
	}

}
