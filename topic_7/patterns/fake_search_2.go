package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result2 string

type Search2 func(query string) Result2

// asyncronous search
func Google2(query string) (results []Result2) {
	c := make(chan Result2)
	go func() { c <- Web2(query) }()
	go func() { c <- Image2(query) }()
	go func() { c <- Video2(query) }()

	// don’t wait for slow servers (Fan-in + timeout for whole conversation)
	timeout := time.After(40 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

var (
	Web2   = fakeSearch2("web")
	Image2 = fakeSearch2("image")
	Video2 = fakeSearch2("video")
)

func fakeSearch2(kind string) Search2 {
	return func(query string) Result2 {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result2(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google2("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
