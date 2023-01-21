package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true


func main()  {
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	// %v prints the value
	result := fmt.Sprintf("x: %v y: %v z: %v", x, y, z)

	fmt.Println("result value is:", result)
	fmt.Println("y value is:", y)
	fmt.Println("z value is:", z)
}