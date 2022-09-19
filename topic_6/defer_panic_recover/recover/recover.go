package main

import "fmt"

func main() {
	b()
	fmt.Println("Returned normally from main.")
}

func b() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in b", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from b.") // this will never be printed
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
