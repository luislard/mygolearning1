package main

import "fmt"

func main() {
	// x := type{values} // COMPOSITE LITERAL
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1:])  // prints the slice from position 1 -> [5,7,8,42]
	fmt.Println(x[1:3]) // prints the slice from position 1 -> [5,7]

}
