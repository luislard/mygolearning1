package main

import "fmt"

func main() {
	a := foo()
	b, c := bar()
	fmt.Println(a)
	fmt.Println(b, c)
}

func foo() int {
	return 5
}

func bar() (int, string) {
	return 5, "hello"
}
