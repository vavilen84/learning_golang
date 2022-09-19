package main

import "fmt"

func main() {
	var m map[int]string
	//m = make(map[int]string)
	m[1] = "string"
	fmt.Printf("%v", m)

}
