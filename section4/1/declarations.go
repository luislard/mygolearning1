package main

import (
	"fmt"
)

var y = "yey"

// DECLARE thre is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS th ZERO VALUE of TYPE int to "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps
var z int

func main() {
	// short declaration operation
	// declare a variable and ASSIGN a value (of certain type)
	x := "ex"
	fmt.Println(x)
	fmt.Println(y)
	foo()
}

func foo()  {
	fmt.Println("again:", y)
	fmt.Println("will print 0 ->", z)
	fmt.Println("z value is:", z)
}