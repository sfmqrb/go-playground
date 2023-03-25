package main

import (
	"fmt"
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

func repeat(
	done <-chan interface{},
	inputs ...interface{},
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		l := len(inputs)
		i := 0
		for {
			select {
			case <-done:
				return
			case valueStream <- inputs[i]:
				i = (i + 1) % l
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

func tee(
	done <-chan interface{},
	in <-chan interface{},
) (_, _ <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})
	go func() {
		defer close(out1)
		defer close(out2)
		for {
			var out1, out2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case out1 <- <-in:
					out1 = nil
				case out2 <- <-in:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
}

func main() {
	done := make(chan interface{})
	defer close(done)
	out1, out2 := tee(done, take(done, repeat(done, 1, 2), 4))
	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
}
