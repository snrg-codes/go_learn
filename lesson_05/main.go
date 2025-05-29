package main

import (
	"fmt"
)

func main() {


	// r := Rectangle{Length: 5, Width: 10}
	// c := Circle{Radius: 7}
	// // PrintShapeInfo(r)
	// // PrintShapeInfo(c)
	// PrintAny(r)
	// PrintAny(c)
	// PrintAny("Hello, World!")
	// PrintAny(42)
	// PrintAny(3.14)
	// PrintAny(true)
	// PrintAny([]int{1, 2, 3})
	// PrintAny(map[string]int{"one": 1, "two": 2})
	r := Rectangle{Length: 5, Width: 10}
	c := Circle{Radius: 7}
	PrintShapeInfo(r)
	PrintShapeInfo(c)
}

func PrintAny(value interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Stringer interface {
	String() string
}

func PrintShapeInfo(s Shape) {
		fmt.Printf("Area: %f\n", s.Area())
		fmt.Printf("Perimeter: %f\n", s.Perimeter())
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle: Length = %f, Width = %f", r.Length, r.Width)
}

type Circle struct {
	Radius float64
}
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle: Radius = %f", c.Radius)
}
