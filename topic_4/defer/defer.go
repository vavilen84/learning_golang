package main

import "fmt"

func main() {
	// LIFO
	defer fmt.Println("!")
	defer fmt.Println("world")

	fmt.Println("hello")
}
