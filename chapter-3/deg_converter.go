package main

import "fmt"

func main() {

	var convType byte

	fmt.Println("Please Choose:\n a: for F to C\n b: fot C to F")
	//fmt.Scanf("%c", &convType)
	convType = 'a'
	if convType == 'b' { //C to F
		var c float32
		fmt.Println("Enter the degree in C:")
		fmt.Scanf("%f", &c)
		f := c*(9/5) + 32
		fmt.Printf("The degree in F is %f\n", f)
	} else if convType == 'a' { //F to C
		var f float32
		fmt.Println("Enter the degree in F:")
		fmt.Scanf("%f", &f)
		c := (f - 32) * (5 / 9)
		fmt.Printf("The degree in C is %f\n", c)
	} else {
		fmt.Println("bad choice!")
	}
}
