package main

import "fmt"

type I interface {
	M()
}

type test struct {
	string2 string
}

func main() {
	var i I
	i = &test{
		string2: "k",
	}
	describe(i)
	i.M()
}

func (t *test) M() {
	fmt.Println("hello")
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
