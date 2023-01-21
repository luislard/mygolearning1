package main

import "fmt"

/*
Using a composite literal

	create a slice which holds the following values
	[42, 43, 44, 45, 46, 47, 48, 49, 50, 51]

append to that this value
52
print out the slice
in One statement append to that slice these
53
54
55
print out
append to the slice this slice
y := []int{56,57,58,59,60}
print out
*/
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

}
