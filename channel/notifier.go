package main

import (
	"fmt"
	"sync"
)

func main() {
	var notifier chan interface{}
	notifier = make(chan interface{})

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-notifier
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines")
	close(notifier)
	wg.Wait()
}
