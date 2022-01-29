package main

import (
	"fmt"
	"math"
)

func main() {

	var num1 int = 8
	var num2 int = 4
	var num3 float64 = 10.0

	answer := num1 + num2
	answer = num1 - num2
	answer = num1 * num2
	answer = num1 % num2
	answer = num1 / num2

	answer1 := math.Sqrt(num3)

	fmt.Printf("%04d\n", answer)
	fmt.Printf("%.2f\n", answer1)

}
