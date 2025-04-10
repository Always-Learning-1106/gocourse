package basics

import "fmt"

func main() {
	// var arrayName [size]elementType
	// var numbers [5]int
	// fmt.Println(numbers)
	//
	// numbers[4] = 20
	// fmt.Println(numbers)
	// numbers[0] = 9
	// fmt.Println(numbers)
	//
	// fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	// fmt.Println("Fruits array:", fruits)
	// fmt.Println("Third element:", fruits[2])
	// originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray
	// copiedArray[0] = 100
	// fmt.Println("Original Array", originalArray)
	// fmt.Println("Copied Array:", copiedArray)
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println("Element at index", i, ":", numbers[i])
	// }
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }
	// a, _ := someFunction()
	// fmt.Println(a)
	// b := 2
	// _ = b
	// fmt.Println("The length of the numbers array is", len(numbers))
	// array1 := [3]int{1, 2, 3}
	// array2 := [3]int{1, 2, 3}
	// fmt.Println("Array1 is equal to Array2", array1 == array2)
	// var matrix [4][3]int = [4][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// 	{10, 11, 12},
	// }
	// fmt.Println(matrix)
	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int
	copiedArray = &originalArray
	copiedArray[0] = 100

	copiedArray[0] = 100
	fmt.Println("originalArray:", originalArray)
	fmt.Println("copiedArray:", *copiedArray)
}
