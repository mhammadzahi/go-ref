package main

import (
	"fmt"
)

func main_jg() {
	var (
		aa = 5
		bb = 7
		cc = 9
	)
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	var a int32
	var b int32
	fmt.Println("Enter a: ")
	fmt.Scanf("%d", &a)

	fmt.Println("Enter b: ")
	fmt.Scanf("%d", &b)

	fmt.Printf("Sum = %d", a+b)

}
