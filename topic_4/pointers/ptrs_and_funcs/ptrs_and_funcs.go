package main

import "fmt"

type A struct {
	Data int
}

func modValue(a A) {
	a.Data = 1
}

func modPtr(a *A) {
	a.Data = 1
}

func main() {
	a := A{}
	fmt.Printf("initial: %v \r\n", a)
	modValue(a)
	fmt.Printf("after modValue: %v \r\n", a)
	modPtr(&a)
	fmt.Printf("after modPtr: %v \r\n", a)
}
