//package main
//
//import "fmt"
//
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
//	fmt.Printf("Type: %T Value: %v\n", false, false)
//	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
//	fmt.Printf("Type: %T Value: %v\n", z, z)
//	i := 42
//	f := float64(i)
//	u := uint(f)
//	fmt.Println(u)
//}


//package main
//
//import "fmt"
//
//const (
//	Big float64 = 1 << 511 + 1 << 511
//)
//
//func needFloat(x float64) float64 {
//	return x * 0.1
//}
//
//func main() {
//	fmt.Println(needFloat(Big))
//}


package main

import "fmt"

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
		if x % base == 0 {
			fmt.Println(x / base)
		}
	}
}
