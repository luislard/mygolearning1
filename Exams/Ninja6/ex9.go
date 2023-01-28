package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("After Setup")
	}

	setup(a)
}

func setup(cb func()) {
	fmt.Println("Starting Setup...")
	cb()
}
