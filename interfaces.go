package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
}

type Square struct {
	size float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64  {
	return s.size * s.size
}

func (c Circle) area() float64  {
	return math.Pi * c.radius * c.radius
}

func measure(g Geometry) float64  {
	return g.area()
}

func main()  {
	c := Circle{5}
	fmt.Println("Circle Area(5)", measure(c))

	s := Square{5}
	fmt.Println("Square Area(5)", measure(s))
}