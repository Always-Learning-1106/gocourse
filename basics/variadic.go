package main

import "fmt"

func main() {
	// Ellipsis
	// func functionName(param1 type1,param2 typ2, param3 ...type3)returnType{
	// function body
	// }
	fmt.Println("Sum of 1,2,3 = ", sum(1, 2, 3))
	numbers := []int{1, 2, 3, 4}
	fmt.Println("sum of numbers slice", sum(numbers...))
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
