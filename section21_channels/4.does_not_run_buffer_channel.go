package main

import (
	"fmt"
)

func main() {
	// there is a buffer of 1 in this channel
	c := make(chan int, 1)

	c <- 42
	c <- 43

	fmt.Println(<-c)

}
