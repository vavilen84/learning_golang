package main

import "fmt"

type A struct {
	Data int
}

type B struct {
	Data int
}

func (a *A) mod() {
	a.Data = 1
}

func (b B) mod() {
	b.Data = 1
}

func main() {
	a := A{}
	b := B{}
	a.mod()
	b.mod()
	fmt.Printf("%v \r\n", a)
	fmt.Printf("%v \r\n", b)
}
