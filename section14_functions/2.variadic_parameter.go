package main

import "fmt"

func main() {
	total := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total is:", total)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i, "we are adding the value", v, "and now the sum is", sum)

	}
	return sum
}
