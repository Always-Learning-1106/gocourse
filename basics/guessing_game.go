package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	// fmt.Println(source)
	random := rand.New(source)
	fmt.Println(random)
	//Generate a random number between 1 and 100
	target := random.Intn(100) + 1
	fmt.Println(target)

}
