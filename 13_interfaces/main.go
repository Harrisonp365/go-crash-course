package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rect struct {
	width, height float64
}

func(circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func(rect Rect) area() float64 {
	return rect.height * rect.width
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y:0, radius: 5}
	rectangle := Rect{width: 10, height: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}