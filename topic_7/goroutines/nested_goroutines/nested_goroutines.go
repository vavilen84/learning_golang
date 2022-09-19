package main

import (
	"fmt"
	"time"
)

// if main finishes earlier than goroutines - goroutiness will be stoped
// but if goroutine finishes earlier than sub-goroutines - sub-goroutines will not be stopped
func main() {
	fmt.Println("start")
	go func() {
		fmt.Println("gmain start")
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println("g1")
		}()
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println("g2")
		}()
		fmt.Println("gmain end")
	}()
	fmt.Println("sleep")
	time.Sleep(5 * time.Second)
	fmt.Println("the end")
}
