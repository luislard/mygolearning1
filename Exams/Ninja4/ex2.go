package main

import "fmt"

/*
Using a composite literal

	create an slice which holds 5 values of type int
	assign values to each index position

Range over the array and print the values out
Using format printing

	print out the type of the slice
*/
func main() {
	x := []int{1, 2, 3, 4, 5}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
