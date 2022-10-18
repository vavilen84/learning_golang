package main

import (
	"fmt"
	"sync"
)

func HelloWorld(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello World!")
}

func main() {
	// TODO: try to remove wg
	var wg sync.WaitGroup
	fmt.Println("Start ...")
	wg.Add(1)
	go HelloWorld(&wg)
	wg.Wait()
	fmt.Println("End")
}
