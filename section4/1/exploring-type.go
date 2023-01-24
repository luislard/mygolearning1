package main

import (
	"fmt"
)

var y = 42
var z string = "Shanken, not stirred"
var a string = `James said: 
"Shanken, not stirred"`

func main()  {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}