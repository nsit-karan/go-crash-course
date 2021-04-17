package main

import (
	"fmt"
	"math"
)

// go figures out that a particular
// struct has a function with the same signature
// as the interface and can use it like a base
// class
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{10, 5}
	fmt.Printf("Circle area is : %f, square area is : %f", getArea(circle), getArea(rectangle))
	fmt.Println("Done!!")
}
