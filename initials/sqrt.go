package main

import (
	"fmt"
	"math"
)

const precision = 1e-20

func sqrtNewtonsMethod(mainNum float64) (float64, int64) {
	return newtonsMethod(1, mainNum, 0)
}

func newtonsMethod(x float64, mainNumber float64, numOfRec int64) (float64, int64) {
	numOfRec += 1
	if precision > math.Abs(math.Pow(x, 2)-mainNumber) {
		return x, numOfRec
	}
	var newX = (1.0 / 2) * (x + mainNumber/x)
	return newtonsMethod(newX, mainNumber, numOfRec)
}
func sqrtNewtonsMethodUsingWhile(mainNumber float64) (float64, int64) {
	x := 1.0
	var numOfRec int64 = 1
	for math.Abs(math.Pow(x, 2)-mainNumber) > precision {
		x = (1.0 / 2) * (x + mainNumber/x)
		numOfRec += 1
	}
	return x, numOfRec
}

func main() {
	x, numOfRec := sqrtNewtonsMethod(81 * 9)
	fmt.Println(x, numOfRec)

	x, numOfRec = sqrtNewtonsMethodUsingWhile(81 * 9)
	fmt.Println(x, numOfRec)
}
