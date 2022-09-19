package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring6("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("You're too slow.")
			return
		}
	}
}

func boring6(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}
