package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i == 1 {
			continue
		}
		if i == 3 {
			break
		}
		fmt.Printf("simple loop: %v \r\n", i)
	}

	// The init and post statements are optional.
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("sum:", sum)

	// infinite loop
	for {
	}
}
