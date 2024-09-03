package main

import (
	"fmt"
)

type shape interface{ area() float64 }
type rect struct{ length, width float64 }

func (r *rect) area() float64 { return r.length * r.width }

type circle struct{ x, y, radius float64 }

func (circle, circle) area() float64 {
	return math.pi * circle.radius * circle.radius
}
func (rect rect) area() float64 {
	return rect.width * rect.height
}
func getArea(shape shape) float64 {
	return shape.area()
}
func main() {
	circle := circle{x: 0, y: 0, radius: 5}
	rectangle := rect{length: 10, width: 5}
	fmt.Println("Circle Area: ", getArea(&circle))
	fmt.Printf("Rectangle Area: %.2f\n", getArea(&rectangle))
}
