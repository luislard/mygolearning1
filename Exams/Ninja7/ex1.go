package main

import "fmt"

/*
*
Create a value and assign it to a variable.
Print the address of that value
*/
func main() {
	x := 23
	fmt.Println(&x)
}
