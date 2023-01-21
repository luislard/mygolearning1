package main

import "fmt"

func main() {
	// x := type{values} // COMPOSITE LITERAL
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x) // prints -> [4 5 7 8 42]
	// print the length of slice
	fmt.Println(len(x)) // prints -> 5

	// access a value
	fmt.Println(x[2]) // prints -> 7 (3rd position in zero based index)

	// using range to loop through a slice
	for i, v := range x {
		fmt.Println(i, v)
	}

}

// a SLICE allows you t group together VALUES of the same TYPE
