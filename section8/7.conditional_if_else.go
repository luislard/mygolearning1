package main

import "fmt"

func main() {
	x := 42

	// if else
	if x == 40 {
		fmt.Println("value is 40")
	} else {
		fmt.Println("our value is not 40")
	}

	// if else if
	if x == 40 {
		fmt.Println("value is 40")
	} else if (x == 41) {
		fmt.Println("our value is 41")
	} else {
		fmt.Println("our value is not 40 or 41")
	}

}
