package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println("All numbers sum", s)

	s2 := even(sum, ii...)
	fmt.Println("Even sum", s2)
}

func sum(intSlice ...int) int {
	total := 0
	for _, v := range intSlice {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if (v%2 == 0) {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
