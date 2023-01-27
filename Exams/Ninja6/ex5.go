package main

import (
	"fmt"
	"math"
)

func main() {
	square := square{
		sideLong: 5,
	}
	circle := circle{
		radius: 1,
	}
	shapeAreaPrinter(square)
	shapeAreaPrinter(circle)
}

func shapeAreaPrinter(s shape) {
	a := s.calcArea()
	fmt.Println(a)
}

type square struct {
	sideLong int
}

type circle struct {
	radius int
}

type shape interface {
	calcArea() float32
}

func (s square) calcArea() float32 {
	return float32(s.sideLong * s.sideLong)
}

func (c circle) calcArea() float32 {
	return math.Pi * float32((c.radius * c.radius))
}
