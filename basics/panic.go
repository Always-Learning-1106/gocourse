package main

import "fmt"

func main() {
	// panic(interface{})
	process(-20)
}

func process(input int) {
	if input < 0 {
		panic("input must be greater than -1")
	}
	fmt.Println("Processing input", input)
}
