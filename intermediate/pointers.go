package main

import "fmt"

func main() {
	var ptr *int
	var a int = 10
	c := 10.0
	fmt.Println(c)
	ptr = &a
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(ptr)
	fmt.Println(&ptr)
	fmt.Println(*ptr)
}
