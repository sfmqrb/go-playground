package main

import "fmt"

type Rectangle struct {
	Width  int8
	Height int8
}

func main() {
	var rec = Rectangle{Width: 1, Height: 3}
	fmt.Println(&rec.Width)
	fmt.Println(&rec.Height)
	x := &rec
	fmt.Println(x.Width)
}
