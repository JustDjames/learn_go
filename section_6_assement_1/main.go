package main

import (
	"fmt"
	"math"
)

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{
		sideLength: 12.5,
	}

	t := triangle{
		height: 10.5,
		base:   20,
	}

	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
