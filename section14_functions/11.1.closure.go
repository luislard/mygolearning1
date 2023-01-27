package main

import "fmt"

// package level scope
var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)

}

func foo() {
	fmt.Println("Hello")
	x++
}
