package main

import "fmt"

func main() {
	num_writes := 5 // also test large numbers like 10000
	chanOwner := func() <-chan int {
		resultStream := make(chan int, num_writes)
		go func() {
			defer close(resultStream)
			for i := 0; i <= num_writes; i++ {
				fmt.Printf("Sent: %d\n", i)
				resultStream <- i
			}
		}()
		return resultStream
	}
	resultStream := chanOwner()

	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
