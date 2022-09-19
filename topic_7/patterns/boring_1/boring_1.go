package main

import (
	"fmt"
	"time"
)

//	we will have no result. The reason is that main goroutine (‘main’ function call) finished earlier,
//
// then we had output from ‘go boring’ goroutine call. The end of main goroutine stops all sub-goroutines.
func main() {
	go boring("boring!")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}
