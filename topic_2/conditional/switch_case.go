package main

import "fmt"

func main() {
	a := 1

	switch a {
	case 1:
		fmt.Println("a equals 1")
		fallthrough
	case 2:
		fmt.Println("a equals 2")
	case 3:
		fmt.Println("a equals 3")
	default:
		fmt.Printf("a equals %v", a)
	}
}
