package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	for k := 0; k < len(a); k++ {
		fmt.Printf("for loop for slice: %v \r\n", a[k])
	}

	for k, v := range a {
		fmt.Printf("key+value: key: %v; value: %v \r\n", k, v)
	}

	for _, v := range a {
		fmt.Printf("skip key: value: %v \r\n", v)
	}

	for k := range a {
		fmt.Printf("only key: %v \r\n", k)
	}
}
