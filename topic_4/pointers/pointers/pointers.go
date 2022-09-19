package main

import "fmt"

func main() {
	a := 1
	b := &a
	*b = 3
	fmt.Println(a)
}
