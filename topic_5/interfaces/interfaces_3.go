package main

import (
	"fmt"
	"math"
)

type Debuger interface {
	Println()
}

type Data struct {
	Data string
}

func (t *Data) Println() {
	fmt.Println(t.Data)
}

type F float64

func (f F) Println() {
	fmt.Println(f)
}

func main() {
	var i Debuger

	i = &Data{"Hello"}
	describe(i)
	i.Println()

	i = F(math.Pi)
	describe(i)
	i.Println()
}

func describe(i Debuger) {
	fmt.Printf("(%v, %T)\n", i, i)
}
