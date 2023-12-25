package main

import (
	"fmt"
	"test/mathoperations"
)

const (
	sum = "sum"
	div = "div"
)

func main() {
	// fmt.Println("hello world")

	var input1, input2 float64
	var operation string
	// input := 10
	input1 = 10
	input2 = 20

	operation = div

	switch operation {
	case sum:
		fmt.Println(mathoperations.Sum(input1, input2))
	case div:
		fmt.Println(mathoperations.Div(input1, input2))
	default:
		fmt.Printf("unkown operation: %s\n", operation)
	}

}
