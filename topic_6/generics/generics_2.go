package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	list := List[int]{
		next: &List[int]{},
		val:  1,
	}
	fmt.Printf("%v \r\n", list)
}
