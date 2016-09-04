package main

import (
	"math"
	"fmt"
)

func main() {
	rect := Rectangle{20, 50}
	circ := Circle{4}

	fmt.Println("Rect area", getArea(rect))
	fmt.Println("Circle area", getArea(circ))
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// method calls appropriate implementation
// via composition
func getArea(shape Shape) float64 {
	return shape.area()
}
