package main

import "fmt"

func main() {

	numbers := []int{2, 3, 4}

	sum := foo(numbers...)
	sum2 := bar(numbers)

	fmt.Println(sum)
	fmt.Println(sum2)
}

func foo(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func bar(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
