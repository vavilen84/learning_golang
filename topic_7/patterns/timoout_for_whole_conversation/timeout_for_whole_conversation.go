package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring7("Joe")
	timeout := time.After(1 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
}

func boring7(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}
