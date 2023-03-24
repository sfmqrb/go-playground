package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int, 2)
	close(ch)

	var ch2 = make(chan int, 2)
	go func() {
		i := 1
		fmt.Println(i)
		ch2 <- i
		close(ch2)
		fmt.Println("after close", i)

	}()

	time.Sleep(time.Second)

	for i := 0; i < 222; i++ {
		select {
		case x, ok := <-ch:
			fmt.Println("closed", x, ok)
		case x, ok := <-ch2:
			fmt.Println("open", x, ok)
			if ok {
				// even though channel is closed
				// (because next line panics) we
				// get true from ok
				// this indicates that closeness of
				// channels are related to the input
				ch <- 1000
			}

		}

	}

}
