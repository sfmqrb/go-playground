package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	str := ""
	for _, c := range s {
		if c == ' ' {
			m[str] += 1
			str = ""
		} else {
			str += string(c)
		}
	}
	if len(str) > 0 {
		m[str] += 1
	}
	delete(m, "")
	return m
}

func main() {
	wc.Test(WordCount)
}
