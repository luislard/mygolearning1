package main

import "fmt"

func main() {
	counter := 10
	inc := foo()
	for i := 0; i < 10; i++ {
		counter++
		inc()
		fmt.Println("External counter is", counter)
	}

}

func foo() func() {
	counter := 0
	return func() {
		for i := 0; i < 3; i++ {
			counter++
			fmt.Println("Internal counter is:", counter)
		}
	}

}
