package main

import (
	"fmt"
)

var y string
var z int

func main()  {
	fmt.Println("value of y ->", y, "<-")
	fmt.Printf("%T\n", y)

	fmt.Println("value of z ->", z, "<-")
	fmt.Printf("%T\n", z)
}