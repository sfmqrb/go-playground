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

func main() {
	x, numOfRec := sqrtNewtonsMethod(81 * 9)
	fmt.Println(x, numOfRec)
}
