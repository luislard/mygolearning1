package main

import "fmt"

func main() {

	x := bar()
	fmt.Printf("%T\n", x)

	fmt.Println(x())

}

func bar() func() int {
	return func() int {
		return 451
	}
}
