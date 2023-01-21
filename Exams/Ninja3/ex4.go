package main

import "fmt"

/*
create a for loop using this syntax
for {}
have it print out the years you have been alive
*/

func main() {
	birdthDate := 1985
	currentYear := 2023
	x := birdthDate
	for {
		if x > currentYear {
			break
		}
		fmt.Println(x)
		x++
	}
}
