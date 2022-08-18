package main

import (
	"fmt"
)

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

	// Like for, the if statement can start with a short statement to execute before the condition.
	if v := 1; v < 3 {
		fmt.Println("v is less than 3")
	} else {
		fmt.Println("v is greater than 3")
	}
}
