package main

import "fmt"

func main() {
	a := 1
	b := 2

	if a == 1 {
		fmt.Println("a equal 1")
	} else {
		fmt.Println("a not equal 1")
	}

	// parentheses are optional
	if a == 1 {
		fmt.Println("a equal 1")
	} else {
		fmt.Println("a not equal 1")
	}

	if a == 1 && b == 2 {
		fmt.Println("a equal 1 & b equal 2")
	}
}
