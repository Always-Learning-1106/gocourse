package main

import "fmt"

func main() {
	process()
}

func process() {
	defer fmt.Println("Deferred statement that was executed")
	defer fmt.Println("Second Deferred statement that was executed")
	defer fmt.Println("Third Deferred statement that was executed")
	fmt.Println("Normal execution statement")
}
