package main

import "fmt"

func aggregator() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := aggregator(), aggregator()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	newAgg := aggregator()
	fmt.Println(newAgg(10), newAgg(20), newAgg(40))
}
