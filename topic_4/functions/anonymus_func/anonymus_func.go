package main

import (
	"fmt"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return x + y
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
}
