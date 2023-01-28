package main

import "fmt"

func main() {
	printHello := func() {
		fmt.Println("Hello ex7")
	}

	printHello()
	fmt.Printf("%T\n", printHello)

}
