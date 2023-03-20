package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	read := func() {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("Reading: ", count)
	}

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Println("Writing: ", count)
	}

	var readWriteGroup sync.WaitGroup

	threshold := 0.9
	for i := 0; i <= 1000; i++ {
		readWriteGroup.Add(1)
		go func() {
			defer readWriteGroup.Done()

			if rand.Float64() > threshold {
				increment()
			} else {
				read()
			}
		}()
	}
	readWriteGroup.Wait()
	fmt.Println("end")
}
