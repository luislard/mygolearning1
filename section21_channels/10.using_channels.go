package main

import "fmt"

func main() {
	c := make(chan int)
	go foo(c)
	bar(c)

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// send
func bar(c <-chan int) {
	fmt.Println(<-c)
}
