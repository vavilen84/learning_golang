package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result3 string
type Search3 func(query string) Result3

var (
	Web21   = fakeSearch3("web1")
	Web22   = fakeSearch3("web2")
	Image21 = fakeSearch3("image1")
	Image22 = fakeSearch3("image2")
	Video21 = fakeSearch3("video1")
	Video22 = fakeSearch3("video2")
)

func Google3(query string) (results []Result3) {
	c := make(chan Result3)
	go func() { c <- First3(query, Web21, Web22) }()
	go func() { c <- First3(query, Image21, Image22) }()
	go func() { c <- First3(query, Video21, Video22) }()
	timeout := time.After(80 * time.Millisecond)
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

func First3(query string, replicas ...Search3) Result3 {
	c := make(chan Result3)
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func fakeSearch3(kind string) Search3 {
	return func(query string) Result3 {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result3(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google3("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
