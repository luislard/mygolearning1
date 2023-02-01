package main

import (
	"encoding/json"
	"fmt"
)

// you can specify how to encode / decode by using "field tags" in the properties see:
// https://pkg.go.dev/encoding/json@go1.19.5#Marshal:~:text=Examples%20of%20struct%20field%20tags%20and%20their%20meanings%3A

// A cool site to get the struct definition from a json is https://mholt.github.io/json-to-go/
type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nall of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
