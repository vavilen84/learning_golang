package main

import "fmt"

func a() {
	defer fmt.Println("defer from a")

	fmt.Println("hello from a")
	var a *int
	fmt.Printf("%v", *a) // panic
	//panic("error!")

	fmt.Println("this will not be printed")
}

func f() {
	defer fmt.Println("defer from f")

	fmt.Println("hello from f")
	a()

	fmt.Println("this will not be printed")
}

func main() {
	defer fmt.Println("defer from main")

	fmt.Println("hello from main")
	f()

	fmt.Println("this will not be printed")
}
