package main

import "fmt"

func main() {
	// an array is a NUMBERED sequence of elements of a single type, called the element type.
	// the following means that we are defining a variable x that is of type array of int of size 5
	var x [5]int
	fmt.Println(x) // -> prints [0 0 0 0 0]

	// lets assign 1 value
	x[3] = 42
	fmt.Println(x) // -> prints [0 0 0 42 0]

	// we can print the length
	fmt.Println(len(x)) // -> prints 5

}
