package main

import "fmt"

/*
*

chan T          // can be used to send and receive values of type T
chan<- float64  // can only be used to send float64s
<-chan int      // can only be used to receive ints
*/
func main() {

	c := make(<-chan int, 2) // receive only channel

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("----------------")
	fmt.Printf("%T\n", c)
	/**
	.\9.directional_channels_receive_only_channel.go:16:2: invalid operation: cannot send to receive-only channel c (variable of type <-chan int)
	.\9.directional_channels_receive_only_channel.go:17:2: invalid operation: cannot send to receive-only channel c (variable of type <-chan int)

	*/
}
