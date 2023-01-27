package main

import "fmt"

func main() {
	foo()
	bar()
	defer foo()
	bar()
	bar()
	fmt.Println("End of Main")
}
func foo() {
	defer fmt.Println("I am Defered in Foo Fooooo")
	fmt.Println("I am the end of Fooooo")
}

func bar() {
	fmt.Println("I am Baaaaaaar")
}
