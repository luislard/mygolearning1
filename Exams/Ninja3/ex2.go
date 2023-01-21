package main

import "fmt"

/*
Print every run code point of the uppercase alphabet three times
should look like:
1

	U+0041 'A'
	U+0041 'A'
	U+0041 'A'

2

	U+0042 'B'
	U+0042 'B'
	U+0042 'B'
*/
func main() {
	// uppercase alpahbet https://en.wikipedia.org/wiki/ASCII#Printable_characters
	// A is 65, Z is 90

	count := 1
	for i := 65; i <= 90; i++ {
		fmt.Println(count)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
		count++
	}
}
