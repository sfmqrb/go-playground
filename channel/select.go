package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			time.Sleep(3 * time.Second)
			x, y = y, x+y
			fmt.Println(time.Now())
		case <-quit:
			fmt.Println("quit")
			return

		}

	}

}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(5 * time.Second)
			fmt.Println(<-c)
			fmt.Println(time.Now())
		}
		quit <- 0

	}()
	fibonacci(c, quit)

}
