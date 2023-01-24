package main

import (
	"fmt"
)

type Hotdog string
var x Hotdog 

func main()  {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)

	x = "The Hotdog"
	fmt.Printf("%v\n", x)

}