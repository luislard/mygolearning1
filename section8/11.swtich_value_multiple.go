package main

import "fmt"

func main() {
	x := "Bond"
	switch x {
	case "Bond", "Dua":
		fmt.Println("A big star")
	default:
		fmt.Println("this is default")

	}

}
