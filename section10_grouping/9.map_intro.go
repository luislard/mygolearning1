package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	// print the entire map
	fmt.Println(m) // prints -> map[James:32 Miss Moneypenny:27]

	// print a specific key
	fmt.Println(m["James"]) // 32

	// what happens if we try to access a value that is not in the map?
	fmt.Println(m["bbb"]) // 0 -> whaaaaat!? yes, this could be a problem
	// there is a way to test if the key exists
	v, ok := m["bbb"]

	fmt.Println(v)
	fmt.Println(ok)

	// HOW TO USE IT IN CONDITIONALS?
	// using this ideomatic chunk of code
	if v, ok := m["bbb"]; ok {
		fmt.Println("this is the IF PRINT for bbb", v) // This will not print since ok == false
	}

	// Now lets try with James
	if v, ok := m["James"]; ok {
		fmt.Println("this is the IF PRINT for James", v) // This will print
	}

}
