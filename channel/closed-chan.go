package main

import "fmt"

func main() {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		stringStream <- "Hello channels!"
	}()

	salutation, ok := <-stringStream
	fmt.Printf("(%v): `%v`\n", ok, salutation)

	salutation, ok = <-stringStream
	fmt.Printf("(%v): `%v`\n", ok, salutation)
}
