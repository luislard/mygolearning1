package main

import "fmt"

func main() {
	withGreting := greetingConfig("Hello Mr:")
	withGreting("Luis")
}

func greetingConfig(greeting string) func(a string) {
	return func(name string) {
		fmt.Println(greeting, name)
	}
}
