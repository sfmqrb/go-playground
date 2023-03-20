// An implementation of Conway's Game of Life.
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type Ground struct {
	space [][]bool
	w, h  int
}

func NewGround(w, h int) *Ground {
	space := make([][]bool, h)
	for i := range space {
		space[i] = make([]bool, w)
	}
	return &Ground{space: space, w: w, h: h}
}

func (g *Ground) Set(x, y int, b bool) {
	g.space[y][x] = b
}

func (g *Ground) GetValidXY(x, y int) (valid_x, valid_y int) {
	x += g.w
	x %= g.w
	y += g.h
	y %= g.h
	return x, y
}

func (g *Ground) Alive(x, y int) bool {
	return g.space[y][x]
}

func (g *Ground) Next(x, y int) bool {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			valid_x, valid_y := g.GetValidXY(x+i, y+j)
			if (j != 0 || i != 0) && g.Alive(valid_x, valid_y) {
				alive++
			}
		}
	}
	return alive == 3 || alive == 2 && g.Alive(x, y)
}

type Life struct {
	a, b *Ground
	w, h int
}

func NewLife(w, h int) *Life {
	a := NewGround(w, h)
	for i := 0; i < (w * h / 4); i++ {
		a.Set(rand.Intn(w), rand.Intn(h), true)
	}
	return &Life{
		a: a, b: NewGround(w, h),
		w: w, h: h,
	}
}

func (l *Life) Step() {
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			l.b.Set(x, y, l.a.Next(x, y))
		}
	}
	l.a, l.b = l.b, l.a
}

func (l *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			b := byte(' ')
			if l.a.Alive(x, y) {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func main() {
	l := NewLife(40, 15)
	for {
		l.Step()
		fmt.Print("\x0c", l)
		time.Sleep(time.Second / 30)
	}
}
