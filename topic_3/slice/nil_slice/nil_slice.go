package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Printf("%v \r\n", s)
		fmt.Println("nil!")
	}

	s[0] = 1
	s = append(s, 1)
	if s != nil {
		fmt.Printf("%v \r\n", s)
		fmt.Println("not nil!")
	}
}
