package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	a := make(chan string)
	go func() {
		a <- "Hello World!"
		wg.Done()
	}()
	go func() {
		fmt.Println(<-a)
		wg.Done()
	}()
	wg.Wait()
}
