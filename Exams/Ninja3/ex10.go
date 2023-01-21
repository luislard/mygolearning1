package main

import "fmt"

/*
create a program that uses a switch expression specified as a variable
of TYPE string with the IDENTIFIER "favSport"
Persistenly, Patienly, You are bound to succeed
*/

func main() {

	favSport := "Baseball"

	switch favSport {
	case "Soccer":
		fmt.Println("Gooooooal")
	case "Baseball":
		fmt.Println("Home RUn")
	}
}
