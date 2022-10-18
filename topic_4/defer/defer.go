package main

import "fmt"

func b() {
	defer a()
}

func a() {
	defer b()
}

func main() {
	// LIFO
	defer a()
	fmt.Println("hello")
}
