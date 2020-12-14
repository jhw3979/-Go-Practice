package main

import (
	"fmt"
)

func main() {
	var A string = "helloworld"
	var B string = "안녕"

	if A == "helloworld" {
		fmt.Println("helloworld")
	} else {
		fmt.Println("hi")
	}

	fmt.Println(len(A))
	fmt.Println(len(B))
}
