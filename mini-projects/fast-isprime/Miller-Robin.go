package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Test if the number 7919 is prime
	n := big.NewInt(7919)
	if isPrime(n, 10) {
		fmt.Println(n, "is probably prime")
	} else {
		fmt.Println(n, "is definitely composite")
	}
}

// isPrime tests if the given number is prime using the Miller-Rabin primality test
// k is the number of witnesses to use (larger k increases confidence)
func isPrime(n *big.Int, k int) bool {
	// Handle small cases
	if n.Cmp(big.NewInt(2)) < 0 {
		return false
	}
	if n.Cmp(big.NewInt(2)) == 0 || n.Cmp(big.NewInt(3)) == 0 {
		return true
	}
	if n.Bit(0) == 0 {
		return false
	}

	// Write n-1 as d*2^r
	r := 0
	d := new(big.Int).Sub(n, big.NewInt(1))
	for d.Bit(0) == 0 {
		r++
		d.Rsh(d, 1)
	}

	// Perform Miller-Rabin test with k witnesses
	for i := 0; i < k; i++ {
		// Choose a random witness a in the range [2, n-2]
		a := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), new(big.Int).Sub(n, big.NewInt(4)))
		a.Add(a, big.NewInt(2))

		// Compute x = a^d mod n
		x := new(big.Int).Exp(a, d, n)

		if x.Cmp(big.NewInt(1)) == 0 || x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
			// Witness 'a' has passed the test
			continue
		}

		// Repeat r-1 times
		passed := false
		for j := 0; j < r-1; j++ {
			x.Exp(x, big.NewInt(2), n)
			if x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
				// Witness 'a' has passed the test
				passed = true
				break
			}
		}

		if !passed {
			// Witness 'a' has failed the test, n is definitely composite
			return false
		}
	}

	// All witnesses have passed the test, n is probably prime
	return true
}

