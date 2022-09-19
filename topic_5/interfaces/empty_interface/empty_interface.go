package main

import (
	"bytes"
	"fmt"
)

func main() {
	var a interface{}
	a = true
	a = 12.34
	a = "hello"
	a = map[string]int{"one": 1}
	a = new(bytes.Buffer)
	fmt.Printf("%v \r\n", a)
}
