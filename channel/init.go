package main

import "fmt"

func main() {
	fmt.Println("let's test channels")
	// I comment the following lines using
	// https://stackoverflow.com/questions/1676632/whats-a-quick-way-to-comment-uncomment-lines-in-vim
	// Yoohhhaaaa
	//	var dataStream chan interface{}
	//	dataStream = make(chan interface{})

	// intStream := make(chan int)

	stringStream := make(chan string)

	go func() {
		stringStream <- "Hello channels"
	}()
	fmt.Println(<-stringStream)
}
