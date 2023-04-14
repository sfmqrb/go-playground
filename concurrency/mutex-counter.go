package main

import (
	"fmt"
	"sync"
	_ "time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string, wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string, wg *sync.WaitGroup) int {
	defer wg.Done()
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey", &wg)
	}
	wg.Wait()
	wg.Add(1)
	fmt.Println(c.Value("somekey", &wg))
}



