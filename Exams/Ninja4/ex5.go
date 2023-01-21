package main

import "fmt"

/*
Using a composite literal

	create a slice which holds the following values
	[42, 43, 44, 45, 46, 47, 48, 49, 50, 51]

delete whats necessary to get [42,43,44,48,49,50,51]
*/
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)

}
