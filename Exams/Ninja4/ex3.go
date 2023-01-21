package main

import "fmt"

/*
Using a composite literal

	create an slice which holds the following values
	[42, 43, 44, 45, 46, 47, 48, 49, 50, 51]

Create these new slices
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]
*/
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

}
