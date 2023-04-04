package main

import (
	"math/big"
	"testing"
)


func main() {
	// Call the testing package's main function to run the tests
	testing.Main(func(_, _ string) (bool, error) { return true, nil }, []testing.InternalTest{{"TestIsPrime", TestIsPrime}}, nil, nil)
}



func TestIsPrime(t *testing.T) {
	// Test some small prime numbers
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}
	for _, p := range primes {
		n := big.NewInt(int64(p))
		if !isPrime(n, 10) {
			t.Errorf("%d should be prime, but isPrime returned false", p)
		}
	}

	// Test some small composite numbers
	composites := []int{4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20, 21, 22, 24, 25}
	for _, c := range composites {
		n := big.NewInt(int64(c))
		if isPrime(n, 10) {
			t.Errorf("%d should be composite, but isPrime returned true", c)
		}
	}

	// Test some larger prime numbers
	primeStrs := []string{
		"3548689",
		"982451653",
		"19491001",
		"1000000007",
	}
	for _, p := range primeStrs {
		n, ok := new(big.Int).SetString(p, 10)
		if !ok {
			t.Fatalf("Failed to parse %s as a big.Int", p)
		}
		if !isPrime(n, 200) {
			t.Errorf("%s should be prime, but isPrime returned false", p)
		}
	}

	// Test some larger composite numbers
	compositeStrs := []string{
		"2465",
		"998244353",
		"19491003",
		"1000000009",
	}
	for _, c := range compositeStrs {
		n, ok := new(big.Int).SetString(c, 10)
		if !ok {
			t.Fatalf("Failed to parse %s as a big.Int", c)
		}
		if isPrime(n, 200) {
			t.Errorf("%s should be composite, but isPrime returned true", c)
		}
	}
}

