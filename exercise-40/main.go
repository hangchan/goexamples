package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func main() {

	c := circle{
		radius: 10,
	}
	// fmt.Println(c.area(c.radius))

	s := square{
		length: 10,
	}
	// fmt.Println(s.area(s.length))

	info(c)

	info(s)

}

func (c circle) area() float64 {
	area := math.Pi * (c.radius * c.radius)
	return area
}

func (s square) area() float64 {
	area := s.length * s.length
	return area
}

func info(s shape) {
	fmt.Println(s.area())
}
