package main

import "fmt"

type C struct {
	X, Y int
}

var (
	v1 = C{1, 2}  // has type C
	v2 = C{X: 1}  // Y:0 is implicit
	v3 = C{}      // X:0 and Y:0
	p  = &C{1, 2} // has type *C
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
