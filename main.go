package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	tri := triangle{
		height: 20,
		base:   30,
	}
	sqr := square{
		sideLength: 17,
	}

	printArea(tri)
	printArea(sqr)
}

func printArea(s shape) {
	fmt.Println("Area of shape:", s.getArea())
}

func (s square) getArea() float64 {
	area := (s.sideLength * s.sideLength)
	return area
}

func (t triangle) getArea() float64 {
	area := (0.5 * t.height * t.base)
	return area
}
