package main

import "fmt"

/*
add a record to the map
*/
func main() {

	m := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["pippo_pluto"] = []string{`Golang`, `PHP`, `Typescript`}

	for key, val := range m {
		fmt.Println("These are the favorite things of", key)
		for _, thing := range val {
			fmt.Printf("\t%v\n", thing)
		}
	}

}
