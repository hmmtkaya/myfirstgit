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

	operation = sum
	mathoperations.A = "hi"

	switch operation {
	case sum:
		fmt.Println(mathoperations.SumOperation(input1, input2))
	case div:
		fmt.Println(float64(input1)/float64(input2))
	default:
		fmt.Printf("unkown operation: %s\n", operation)
	}

	fmt.Println(mathoperations.A)
	mathoperations.A = "hi"

	fmt.Println(mathoperations.A)


}




