package main

import "fmt"

func main() {
	myStr := "Hello world"
	for i := 0; i < len(myStr); i++ {
		char := myStr[i]
		fmt.Printf("%c\n", char)
	}
}
