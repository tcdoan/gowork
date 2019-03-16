package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	size float64
}

func main() {
	t := triangle{
		base:   10.0,
		height: 20.0,
	}
	s := square{
		size: 10.0,
	}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func (s square) getArea() float64 {
	return s.size * s.size
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
