package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	// print the entire map
	fmt.Println(m) // prints -> map[James:32 Miss Moneypenny:27]

	// delete a key
	delete(m, "James")

	fmt.Println(m) // map[Miss Moneypenny:27]

	// you can also delete not existing keys
	delete(m, "bbb")

	// to delete something that is really there
	// first check that the key exists using the comma ideomatic syntax ", ok"
	if _, ok := m["Miss Moneypenny"]; ok {
		delete(m, "Miss Moneypenny")
	}
	fmt.Println(m) // map[]

}
