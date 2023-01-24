package main

import (
	"fmt"
)

type Hotdog int
var x Hotdog
var y int

func main()  {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Printf("%v\n", x)

//	y := string("Hello")
	y := int(x)

	fmt.Printf("%v\n", y)

}