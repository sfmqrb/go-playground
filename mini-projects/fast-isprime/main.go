package main

import (
	"fmt"
	"math"
	"sync"
)

func isPrimeRelative(m, n int) bool {
	return m % n != 0
}

func checkPrime(m, n int, ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- isPrimeRelative(m, n)
}

func isConcurrentPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}

	sqrtN := int(math.Sqrt(float64(n)))
	ch := make(chan bool)
	var wg sync.WaitGroup
	for i := 3; i <= sqrtN; i += 2 {
		wg.Add(1)
		go checkPrime(n, i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for prime := range ch {
		if !prime {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isConcurrentPrime(13))  // true
	fmt.Println(isConcurrentPrime(16))  // false
	fmt.Println(isConcurrentPrime(97))  // true 
	fmt.Println(isConcurrentPrime(100)) // false
}








