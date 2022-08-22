package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	for k, v := range m {
		fmt.Printf("%v %v \r\n", k, v)
	}

	for k := range m {
		fmt.Printf("%v \r\n", k)
	}
}
