package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a()) // 1
	fmt.Println(a()) // 2
	fmt.Println(a()) // 3
	fmt.Println(b()) // 1
	fmt.Println(b()) // 2

	/**
	Notice that when b() is called the x variable gets created for the first time again, why?
	because the incrementor function was stored in memory twice in different locations
	variable a holds the first incrementor
	variable b holds the second incrementor
	*/

}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x

	}
}
