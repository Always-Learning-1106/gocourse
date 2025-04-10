package basics

import "fmt"

func main() {
	var collect []int

	countdownby := 7
	for i := 100; i > 0; i -= countdownby {
		collect = append(collect, i)
	}
	fmt.Println(collect)
	fmt.Println(collect[7])
}
