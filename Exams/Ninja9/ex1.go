package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("start of main")
	wg.Add(2)
	go foo()
	go bar()
	fmt.Println("end of main")
	wg.Wait()
}

func foo() {
	fmt.Println("I am foo")
	wg.Done()
}

func bar() {
	fmt.Println("I am bar")
	wg.Done()
}