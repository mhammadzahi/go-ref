package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//var convType string

	fmt.Println("Please Choose:\n a: for F to C\n b: fot C to F")
	//fmt.Scanf("%s", &convType)
	reader := bufio.NewReader(os.Stdin)
	convType, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
		return
	}
	//convType = 'a'
	if convType == 'b' { //C to F
		var c float32
		fmt.Println("Enter the degree in C:")
		fmt.Scanf("%f", &c)
		f := c*(9.0/5.0) + 32
		fmt.Printf("The degree in F is %f\n", f)
	} else if convType == 'a' { //F to C
		var f float32
		fmt.Println("Enter the degree in F:")
		fmt.Scanf("%f", &f)
		c := (f - 32) * (5.0 / 9.0)
		fmt.Printf("The degree in C is %f\n", c)
	} else {
		fmt.Println("bad choice!")
	}
}
