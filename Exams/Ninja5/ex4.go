package main

import "fmt"

/*
○ Create an anonymus struc type with a field map

	and another with type slice
*/
func main() {
	testee := struct {
		mymap   map[string]string
		myslice []string
	}{
		mymap: map[string]string{
			"foo": "bar",
		},
		myslice: []string{"hello world"},
	}

	fmt.Println(testee)
	fmt.Println(testee.mymap)
	fmt.Println(testee.myslice)
	fmt.Println(testee.mymap["foo"])

}
