package main

import (
	"fmt"
	"strconv"
	"time"
)

func k() {
	panic(100)
}

func j() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in j", r)
		}
	}()
	go k()
}

func main() {
	j()
	for i := 3; i > 0; i-- {
		fmt.Println("main:", strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
