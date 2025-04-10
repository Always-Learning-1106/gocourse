package basics

import (
	"fmt"
	"slices"
)

func main() {
	// var numbers []int
	// numbers1 := []int{1, 2, 3}
	// numbers2 := []int{9, 8, 7}
	// slice := make([]int, 5)
	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]
	fmt.Println(slice1)
	slice1 = append(slice1, 6, 7)
	fmt.Println(slice1)
	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println(sliceCopy)
	// var nilSlice []int
	for i, v := range slice1 {
		fmt.Println(i, v)
	}
	fmt.Println("Element at index 3 of slice1", slice1[3])
	// slice1[3] = 50
	// fmt.Println("Element at index 3 of slice1", slice1[3])
	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 is equal to sliceCopy")
	}
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}

	}
	fmt.Println(twoD)
}
