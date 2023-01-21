package main

import (
	"fmt"
)

var x int = 1
var y string = "some"
var z bool = true


func main()  {
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	fmt.Println("x value is:", x)
	fmt.Println("y value is:", y)
	fmt.Println("z value is:", z)
}