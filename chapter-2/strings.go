package main

import "fmt"

func main_k() { //rename it
	// myStr := "Hello world"
	// for i := 0; i < len(myStr); i++ {
	// 	char := myStr[i]
	// 	fmt.Printf("%c\n", char)
	// }
	myStr := "hello"

	fmt.Printf("%s\n", myStr)

	fmt.Printf("%d\n", len(myStr))

	fmt.Printf("%c\n", myStr[2])

	myStr2 := myStr + " e"

	fmt.Printf("%s\n", myStr2)
	fmt.Printf("%c\n", "hello, World"[1])
}
