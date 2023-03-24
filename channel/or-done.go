package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	myChan := make(chan interface{})
	go func() {
		//	defer close(myChan)
		for i := 1; i < 10; i++ {
			myChan <- i
		}
	}()

	go func() {
		time.Sleep(time.Second * 3)
		close(done)
	}()
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					fmt.Printf("consume %v\n", v)
					if ok == false {
						return

					}
					select {
					case valStream <- v:
						fmt.Printf("consume second layer select %v\n", v)
					case <-done:
						fmt.Println("in second select")
					}

				}

			}

		}()
		return valStream

	}

	for val := range orDone(done, myChan) {
		// Do something with val
		fmt.Println(val)
	}
	//for val := range myChan {
	//	_ = orDone
	//	// Do something with val
	//	fmt.Println(val)
	//}
}
