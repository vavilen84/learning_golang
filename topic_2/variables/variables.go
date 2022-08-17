package main

import "fmt"

/**
Multi-string comment
*/

func main() {
	// Declaration syntax is:
	// var a type = expression
	var a int8 = 127
	fmt.Printf("declared & assigned var a: %T %v \r\n", a, a)

	// equal
	var aa int8
	aa = 127
	fmt.Printf("declared & assigned var aa: %T %v \r\n", aa, aa)

	// auto-assigned type will be int
	b := 127
	fmt.Printf("auto-assigned int: %T %v \r\n", b, b)

	// auto-assigned type will be float64
	f := 127.1
	fmt.Printf("auto-assigned float: %T %v \r\n", f, f)

	d := int8(127)
	fmt.Printf("int8: %T %v \r\n", d, d)

	var e int
	fmt.Printf("initial var e value is zero: %T %v \r\n", e, e)

	// multiple declaration
	var n1, n2, n3 int
	fmt.Printf("multiple declaration n1: %T %v \r\n", n1, n1)
	fmt.Printf("multiple declaration n2: %T %v \r\n", n2, n2)
	fmt.Printf("multiple declaration n3: %T %v \r\n", n3, n3)

	// multiple assignment
	n4, n5 := int8(1), float32(1)
	fmt.Printf("multiple assignment n4: %T %v \r\n", n4, n4)
	fmt.Printf("multiple assignment n5: %T %v \r\n", n5, n5)

	// Variables in Golang are register dependent
	x := 1
	X := 2
	fmt.Printf("var x: %T %v \r\n", x, x)
	fmt.Printf("var X: %T %v \r\n", X, X)
}
