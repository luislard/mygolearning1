package main

import "fmt"

// is when we assign a func to a variable
func main() {
	f := func(x int) {
		fmt.Println("my first func expression", x)
	}

	f(44)
}
