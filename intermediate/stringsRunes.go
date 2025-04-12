package main

import "fmt"

func main() {
	message := "Hello, Go!"
	message2 := "Hello, \tGo!"
	rawMessage := `Hello\nGo`
	fmt.Println(message2)
	fmt.Println(message)
	fmt.Println(rawMessage)
}
