package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)         //42
	fmt.Println(&a)        // 0xc00001c098 the address in memory
	fmt.Printf("%T\n", a)  // int
	fmt.Printf("%T\n", &a) // *int (a pointer to an int) // & gives you the address

	//var b int = &a // you cant assign *int to int

	b := &a
	fmt.Println(b)   // prints an address
	fmt.Println(*b)  // * gives oyu the value stored at a certain address -> prints th value in the pointer (42)
	fmt.Println(*&a) // prints the value at the address &a  -> 42

	*b = 43        // assign 43 to the value in the pointer b (which is the pointer of a)
	fmt.Println(a) // 43

}
