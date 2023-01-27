package main

import "fmt"

func main() {
	numbers := []int{1, 1, 1, 1}
	//numbers := []int{}
	total := sum(numbers...)
	//total := sum()
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

// the variadic paramater (...T) should be in the last position of the parameters

//func iAmIncorrect(x ...int, s string) {
//	fmt.Println(x)
//	fmt.Println(s)
//}

func iAmCorrect(s string, x ...int) {
	fmt.Println(x)
	fmt.Println(s)
}
