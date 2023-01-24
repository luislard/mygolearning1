package main

import (
	"fmt"
)
	// using iota create constant for the last 4 years

const (
	a = 2017 + iota
	b = 2017 + iota
	c = 2017 + iota
	d = 2017 + iota
)

func main()  {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}

