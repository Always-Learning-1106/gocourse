package basics

import "fmt"

func variable() {
	var age int
	var name string = "John"
	var name1 = "Jane"
	age = 12
	count := 10
	lastName := "Smith"
	total := age + count
	fmt.Println(total)
	fmt.Println(lastName, name, name1)
}
