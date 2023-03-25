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
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}

					select {
					// what happens exactly in the next line?
					// if read is successful but write blocked
					// select continue its work in other cases
					// or block the whole select statement
					case valStream <- v:
						fmt.Println("in valStream <- v", v)
					case <-done:
						fmt.Println("in Done", v)
					default:
						fmt.Println("in default")
					}
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
