package main

import "fmt"

var (
	a int
	b int8
	c int16
	d int32
	e int64
)

func main() {
	// cast from int8 to int16
	b = 127
	c = int16(b)
	fmt.Printf("expected result: %T %v \r\n", c, c)

	// cast from int16 to int8
	c = 128
	b = int8(c)
	fmt.Printf("unexpected result: %T %v \r\n", b, b)

	// cast from int16 (in int8 range) to int8
	c = 127
	b = int8(c)
	fmt.Printf("expected result: %T %v \r\n", b, b)
}
