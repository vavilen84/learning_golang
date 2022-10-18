package main

import "fmt"

func sendToChan(c chan int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("writing %d to chan \r\n", i)
		c <- i
	}

	// TODO: try to comment
	close(c)
}

func main() {
	var c = make(chan int)
	go sendToChan(c)
	for v := range c {
		fmt.Printf("reading %d from chan \r\n", v)
	}

	// or
	//for i := 0; i < 5; i++ {
	//	v := <-c
	//	fmt.Printf("reading %d from chan \r\n", v)
	//}
}
