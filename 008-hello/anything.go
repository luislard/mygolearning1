package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	foo()
	fmt.Println("Something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	n, err := fmt.Println("and we exited!", 42, true)
	fmt.Println("bites written", n)
	fmt.Println("error?", err)
	
}
