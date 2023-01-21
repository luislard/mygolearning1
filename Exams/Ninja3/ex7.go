package main

import "fmt"

/*
create a program that shows the if statement in action
*/

func main() {

	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Printf("The val %v is divisible by 4\n", i)
		}
	}
}
