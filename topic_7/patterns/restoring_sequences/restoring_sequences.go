package main

import (
	"fmt"
)

type Message struct {
	str  string
	wait chan bool
}

func main() {
	c := fanIn2(boring4("Joe"), boring4("Ann"))

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true // continues goroutine inside boring4 func
		msg2.wait <- true // continues goroutine inside boring4 func
	}

	fmt.Println("You're all boring; I'm leaving.")
}

func boring4(msg string) <-chan Message { // Returns receive-only channel of strings.
	c := make(chan Message)

	waitForIt := make(chan bool) // Shared between all messages.

	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {

			c <- Message{fmt.Sprintf("%s: %d", msg, i), waitForIt}

			// if we comment timer - <-waitForIt make sync of both channels
			//time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c // Return the channel to the caller.
}

func fanIn2(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i] // New instance of 'input' for each loop.
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}
