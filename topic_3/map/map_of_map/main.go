package main

import "fmt"

func main() {
	var m map[int]map[int]string
	m = make(map[int]map[int]string)
	//m[1] = make(map[int]string)
	m[1][1] = "string"
	fmt.Printf("%v \r\n", m)
}
