package main

import "fmt"

// iota is often used as ENUM
const (
	ready = iota
	blocked
	inProgress
)

const (
	_ = iota
	published
	unpublished
)

func main() {
	fmt.Printf("ready %T %v \r\n", ready, ready)
	fmt.Printf("blocked %T %v \r\n", blocked, blocked)
	fmt.Printf("inProgress %T %v \r\n", inProgress, inProgress)

	fmt.Printf("published %T %v \r\n", published, published)
	fmt.Printf("unpublished %T %v \r\n", unpublished, unpublished)
}
