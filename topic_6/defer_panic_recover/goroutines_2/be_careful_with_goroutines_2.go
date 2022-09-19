package main

import (
	"fmt"
	"strconv"
	"time"
)

// we should put recover in the panic goroutine
func m() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in m", r)
		}
	}()
	panic(100)
}

func l() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in l", r)
		}
	}()
	go m()
}

func main() {
	l()
	for i := 3; i > 0; i-- {
		fmt.Println("main:", strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
