package main

import (
	"fmt"
	"math/rand"
)

// Probabilistic primality test using Miller-Rabin algorithm
func isPrime(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 || n == 3 {
		return true
	} else if n%2 == 0 {
		return false
	}

	// Find k, q such that n-1 = 2^k * q
	k, q := 0, n-1
	for q%2 == 0 {
		k++
		q /= 2
	}

	// Run the Miller-Rabin test with 10 random witnesses
	for i := 0; i < 10; i++ {
		a := rand.Intn(n-3) + 2 // Random integer in [2, n-2]
		x := powmod(a, q, n)
		if x == 1 || x == n-1 {
			continue
		}
		for j := 0; j < k-1; j++ {
			x = powmod(x, 2, n)
			if x == n-1 {
				break
			}
		}
		if x != n-1 {
			return false
		}
	}

	return true
}

// Calculate (a^b) mod m using binary exponentiation
func powmod(a, b, m int) int {
	result := 1
	for b > 0 {
		if b&1 == 1 {
			result = (result * a) % m
		}
		a = (a * a) % m
		b /= 2
	}
	return result
}

func main() {
	fmt.Println(isPrime(103))
	fmt.Println(isPrime(97))
	fmt.Println(isPrime(19))
	fmt.Println(isPrime(121))
}

