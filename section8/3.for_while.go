package main

import "fmt"

func main() {
	x := 1
	// this is how we express a while loop
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Done.")
}
