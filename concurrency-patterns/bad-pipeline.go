package main

import (
	"fmt"
	"math/rand"
	"time"
)

func take(
	done <-chan interface{},
	valueStream <-chan interface{},
	num int,
) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for counter := 0; counter < num; counter++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func repeatFn(
	done <-chan interface{},
	fn func() interface{},
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func toInt(
	done <-chan interface{},
	valueStream <-chan interface{},
) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for {
			select {
			case <-done:
				return
			case intStream <- (<-valueStream).(int):
			}
		}
	}()
	return intStream
}

func primeFinder(
	done <-chan interface{},
	inputStream <-chan int,
) <-chan interface{} {
	primeStream := make(chan interface{})
	go func() {
		defer close(primeStream)
		for {
			select {
			case <-done:
				return
			case input := <-inputStream:
				isPrime := true
				for j := 2; j <= int(input/2)+1; j++ {
					if input%j == 0 {
						isPrime = false
						break
					}
				}
				if isPrime {
					primeStream <- input
				}
			}
		}
	}()
	return primeStream
}

func main() {
	rand := func() interface{} { return rand.Intn(50000000) }
	done := make(chan interface{})
	defer close(done)
	start := time.Now()
	randIntStream := toInt(done, repeatFn(done, rand))
	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}
