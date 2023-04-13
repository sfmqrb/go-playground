package main

import "fmt"
import "math"

//var isHello bool = true
//var isBye bool
//
//func split(sum int) (x, y int) {
//	x = sum * 4 / 9
//	y = sum - x
//	return
//}
//
//func main() {
//	fmt.Println(split(17))
//	fmt.Println(isBye, isHello)
//	var x int = 19
//	fmt.Println(x)
//}

//package main
//
//import (
//	"fmt"
//	"math/cmplx"
//)
//
//var (
//	MaxInt uint64     = 1<<64 - 1
//	z      complex128 = cmplx.Sqrt(-5 + 12i)
//)
//
//func main() {
//	
//	
//	
//	
//	
//	
//	
//}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	counter := 1

	for ; ; counter *= 2 {
		if counter > 1000 {
			break
		}
		fmt.Println(counter)
	}
	var x = 1
	base := 1_000_000_000
	for {
		x += 1
		if x%base == 0 {
			fmt.Println(x / base)
		}
	}
	var x, y int = 3, 4
	var f float64 = math.Pow((float64(x*x + y*y)), 0.5)
	var z uint = uint(f)
	fmt.Println(x, y, z)
	fmt.Printf("Type: %T Value: %v\n", false, false)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(u)

}
