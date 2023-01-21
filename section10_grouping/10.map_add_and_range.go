package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	// print the entire map
	fmt.Println(m) // prints -> map[James:32 Miss Moneypenny:27]

	// adding a key
	m["Luis"] = 37

	fmt.Println(m) // map[James:32 Luis:37 Miss Moneypenny:27]

	// how to iterate over all keys?
	for key, val := range m {
		fmt.Println(key, val)
	}

}
