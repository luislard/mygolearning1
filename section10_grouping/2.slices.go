package main

import "fmt"

func main() {
	// x := type{values} // COMPOSITE LITERAL
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x) // prints -> [4 5 7 8 42]
}

// a SLICE allows you t group together VALUES of the same TYPE
