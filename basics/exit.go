package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting the main function")
	// exit with status code 1
	os.Exit(1)
	// This will never be executed
	fmt.Println("End of main function")
}
