package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	primes[1] = 1
	fmt.Println(s)
	fmt.Printf("len: %v \r\n", len(s))
	fmt.Printf("cap: %v \r\n", cap(s))

	r := make([]int, 3, 5)
	fmt.Printf("r: %v \r\n", r)
	fmt.Printf("r len: %v \r\n", len(r))
	fmt.Printf("r cap: %v \r\n", cap(r))
}
