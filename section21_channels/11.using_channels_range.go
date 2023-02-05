package main

import "fmt"

func main() {
	c := make(chan int)
	go foo(c)

	for v := range c {

		fmt.Println(v)

	}

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
