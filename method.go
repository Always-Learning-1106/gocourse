package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func main() {
	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Println("Area of rectangle with width 9 and length 10 is", area)
	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area of rectangle with width 9 and length 10 is", area)
}
