package main

import (
	"fmt"
	"time"
)

func main() {
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{}, 4)
		go func() {
			defer close(valStream)
			for {
				select {
				case valStream <- <-c:
					fmt.Println("in valStream <- v")
				case <-done:
					fmt.Println("in Done")
					return
				default:
					fmt.Println("in default")
				}
			}
		}()
		return valStream
	}

	stream := make(chan interface{})
	done := make(chan interface{})
	defer close(done)
	go func() {
		defer close(stream)
		for i := 0; i < 11; i++ {
			stream <- i
		}
	}()

	orDone(done, stream)
	time.Sleep(time.Second)
}
