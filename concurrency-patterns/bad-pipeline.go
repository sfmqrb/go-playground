package main

import (
	"fmt"
	"math/rand"
	"runtime"
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
				// very bad primefinder
				// want to test time and
				// optimize using fan-in/fan-out
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

func fanIn(
	done <-chan interface{},
	channels ...<-chan interface{},
) <-chan interface{} {
	combined := make(chan interface{})
	combiner := func(c <-chan interface{}) {
		for {
			select {
			case <-done:
				return
			case tmp := <-c:
				combined <- tmp
			}
		}
		close(combined)
	}

	for _, channel := range channels {
		go combiner(channel)
	}
	return combined
}

func main() {
	done := make(chan interface{})
	defer close(done)
	start := time.Now()
	rand := func() interface{} { return rand.Intn(50000000) }
	randIntStream := toInt(done, repeatFn(done, rand))

	numFinders := runtime.NumCPU()

	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)

	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}
	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}
