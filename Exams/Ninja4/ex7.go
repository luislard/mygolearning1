package main

import "fmt"

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."
Range over the records, then range over the data in each record
*/
func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	multi := [][]string{x, y}
	fmt.Println(multi)

	for i, sub := range multi {
		fmt.Println("Outer", i)
		for j, subVal := range sub {
			fmt.Println(j, subVal)
		}
	}
}
