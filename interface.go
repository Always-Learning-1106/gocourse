package main

import (
	"fmt"
	"math"
)

type geometry interface {
	Area() float64
	Perim() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Perim() float64 {
	return 2 * (r.height * r.width)
}

func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) diameter() float64 {
	return 2 * c.radius
}

func main() {
	r := Rectangle{width: 3, height: 4}
	s := "this is a string"
	fmt.Println(s, r)
	fmt.Println("This is a test to test my printing skills")
}
