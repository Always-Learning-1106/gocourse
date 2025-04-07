package basics

import (
	"fmt"
	foo "net/http"
)

func basicss() {
	fmt.Println("Hello, GO standard library")
	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return

	}
	defer resp.Body.Close()
	fmt.Println("HTTP Response Status:", resp.Status)
}
