package basics

import "fmt"

func main() {
	// 	count := 5
	// 	for i := 1; i <= count; i++ {
	// 		fmt.Println(i)
	// 	}
	// iteration over collection
	// numbers := []int{1, 2, 3, 4, 5, 6}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Odd:", i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }
	rows := 5
	// Outer loop
	for i := 1; i <= rows; i++ {
		// inner loop for spaces before stars
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		// inner loop for stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
