package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println("anonymous func ran", x)
	}(42)
}
