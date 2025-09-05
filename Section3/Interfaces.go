package Section3

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	peremeter() float64
}

type Circle struct {
	Radius float32
}

type Square struct {
	wideth float32
	height float32
}

func (c Circle) area() float64 {
	return math.Pi * float64(c.Radius)
}

func (s Square) peremeter() float64 {
	return float64(s.height * s.wideth)
}

func Interfaces() {
	c := Circle{Radius: 10}
	s := Square{wideth: 10, height: 20}

	var g geometry

	fmt.Println("Circle Area:", c.area())
	fmt.Println("Square Peremeter:", s.peremeter())
	fmt.Println(g)
}
