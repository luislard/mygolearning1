package main

import "fmt"

/*
*

chan T          // can be used to send and receive values of type T
chan<- float64  // can only be used to send float64s
<-chan int      // can only be used to receive ints
*/
func main() {

	c := make(chan<- int, 2) // send only channel

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("----------------")
	fmt.Printf("%T\n", c)
	/**
.\8.directional_channels_send_only_channel.go:19:16: invalid operation: cannot receive from send-only channel c (variable of type chan<- int)
.\8.directional_channels_send_only_channel.go:20:16: invalid operation: cannot receive from send-only channel c (variable of type chan<- int)

	*/
}
