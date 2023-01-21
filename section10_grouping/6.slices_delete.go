package main

import "fmt"

func main() {
	// x := type{values} // COMPOSITE LITERAL
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{123, 456, 789}
	x = append(x, y...)
	fmt.Println(x) // prints [4 5 7 8 42 77 88 99 1014 123 456 789]

	// =========================
	// Given [4 5 7 8 42 77 88 99 1014 123 456 789]
	// Remove 7 8
	// =========================
	x = append(x[:2], x[4:]...)
	fmt.Println(x) // prints [4 5 42 77 88 99 1014 123 456 789]

}
