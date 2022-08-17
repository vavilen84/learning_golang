package main

import "fmt"

const (
	t = true
	f = false
)

func isEqual(n1, n2 int) bool {
	return n1 == n2
}

func main() {
	res := isEqual(1, 1)
	fmt.Printf("%T %v \r\n", res, res)

	res = isEqual(1, 0)
	fmt.Printf("%T %v \r\n", res, res)
}
