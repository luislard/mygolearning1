package main

import "fmt"

func main() {
	//exampleA()
	exampleB()
}

func exampleA() {
	x := 0
	fooA(x)
	fmt.Println(x) // 0
}

func fooA(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func exampleB() {
	x := 0
	fooB(&x)
	fmt.Println(x) // 43
}

func fooB(y *int) { // receives a pointer to an int (receives an address where an int is located)
	fmt.Println(*y)
	*y = 43 // change the value (*) of the address (y) by assigning 43 to it
	fmt.Println(*y)
}
