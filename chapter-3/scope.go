package main

import (
	"fmt"
)

var x string = "hello"

func main_sco() {
	fmt.Println(f())
}

func f() string {
	res := x + " world"
	return res
}
