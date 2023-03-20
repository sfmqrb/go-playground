package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		fmt.Println("1- here we are!!")
		time.Sleep(3 * time.Second)
		wg.Done()
		fmt.Println("2- here we are!!")
		time.Sleep(3 * time.Second)
		fmt.Println("3- here we are!!")
	}()

	wg.Wait()
	fmt.Println("vim-go")
}
