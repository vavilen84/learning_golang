package main

import "fmt"

func main() {
	a := make([]int, 5)
	ps("a", a)

	b := make([]int, 0, 5)
	ps("b", b)

	c := b[:2]
	ps("c", c)

	d := c[2:5]
	ps("d", d)
}

func ps(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
