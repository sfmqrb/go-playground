package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]Vertex)

	m["someText"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["someText"])
}
