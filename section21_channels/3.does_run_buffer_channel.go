package main

import "fmt"

func main() {
	// there is a buffer of 1 in this channel
	c := make(chan int, 1)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

}
