package main

import "fmt"

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[0] = 9999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423)

	fmt.Println(x)
	fmt.Println(len(x)) // 11
	fmt.Println(cap(x)) // 12

	x = append(x, 3424)
	x = append(x, 3425)

	fmt.Println(x)
	fmt.Println(len(x)) // 13
	fmt.Println(cap(x)) // 24

}
