package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	// This lock is called a multiple readers, single writer.
	// The mechanism allows read operations to be performed concurrently, but write operations are blocking.
	mu sync.RWMutex
	v  map[string]int
}

// But I have to use Lock () for set (), which changes data
func (c *Counter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

// I can safely use RLock () for the counter, since it does not change the data
func (c *Counter) Value(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.v[key]
}

func main() {
	c := Counter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
