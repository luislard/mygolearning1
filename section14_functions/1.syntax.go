package main

import "fmt"

func main() {
	foo()
	bar("Luis")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

// func (r receiver) identifier(parameters) (return(s)) { ... }
func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)

}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(first string, last string) (string, bool) {
	a := fmt.Sprint(first, " ", last, `, says "Hello"`)
	b := false
	return a, b
}
