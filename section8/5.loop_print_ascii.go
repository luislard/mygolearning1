package main

import "fmt"

/*
Challenge:
Print all Ascii characters from 0 - 122
and print them out as text
See: https://en.wikipedia.org/wiki/ASCII#Printable_characters DEC column
See printf formatting in https://pkg.go.dev/fmt#hdr-Printing
*/
func main() {
	fmt.Println("Start")
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
}
