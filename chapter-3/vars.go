package main

import "fmt"

func main_() {
	var varName string
	varName = "hello"
	fmt.Println(varName)
	//--
	var varName2 int = 55
	fmt.Println(varName2)
	//--
	var x string = "hello"
	var y string = "hello "
	fmt.Println(x == y)
	//--
	//go short variable declaration
	varName3 := "abc"
	fmt.Println(varName3)
}
