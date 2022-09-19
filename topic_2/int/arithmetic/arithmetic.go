package main

import "fmt"

func sum(n1, n2, n3 int, n4 int8) int {
	return n1 + n2 + n3 + int(n4)
}

func divInt(n1, n2 int) (res int) {
	res = n1 / n2 // res is also int
	return
}

// rounding https://yourbasic.org/golang/round-float-2-decimal-places/
func divFloat32(n1, n2 float32) (res float32) {
	res = n1 / n2
	return
}

func divFloat64(n1, n2 float64) (res float64) {
	res = n1 / n2
	return
}

func main() {
	resSum := sum(1, 2, 3, 4)
	fmt.Printf("sum: %T %v \r\n", resSum, resSum)

	resInt := divInt(3, 2)
	fmt.Printf("divInt: %T %v \r\n", resInt, resInt)

	resFloat32 := divFloat32(3.345, 1.345345)
	fmt.Printf("divFloat: %T %v \r\n", resFloat32, resFloat32)

	resFloat64 := divFloat64(3.345, 1.345345)
	fmt.Printf("divFloat: %T %v \r\n", resFloat64, resFloat64)
}
