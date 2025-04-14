package main

import "fmt"

func prefixer(op string) func(string) string {
	return func(s string) string {
		return op + s
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}
