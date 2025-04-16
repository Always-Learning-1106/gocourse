package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}
type Shape struct {
	Rectangle
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
	area = rect.Area()
	rect.Scale(2)
	fmt.Println("Area of rectangle with width 9 and length 10 is", area)
	num := Myint(-5)
	num1 := Myint(9)
	fmt.Println(num.isPositive())
	fmt.Println(num1.isPositive())
	var welcome Myint
	welcome2 := Myint(8)
	fmt.Println(welcome.welcomeMessage())
	fmt.Println(welcome2.welcomeMessage())
	s := Shape{Rectangle{width: 12.0, length: 4.0}}
	fmt.Println(s.length, s.width)
	fmt.Println(s.multiplyShape())
}

// Method on a user-defined type
type Myint int

func (m Myint) welcomeMessage() string {
	return "welcome to MyInt Type"
}

func (m Myint) isPositive() bool {
	return m > 0
}

func (s Shape) multiplyShape() float64 {
	return s.length * s.width
}
