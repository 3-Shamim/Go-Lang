package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello")
	fmt.Println(arithmeticOperation(1, 2))
	fmt.Println(arithmeticOperation1(1, 2))

	inner := func() {
		fmt.Println("Inner function")
	}

	inner1 := func(x int) int {
		return x * 3
	}

	inner()

	fmt.Println("Inner func", inner1(3))

	invokedFunc := func(x int) int {
		return x * -1
	}(5)

	fmt.Println("Invoked func", invokedFunc)

	fmt.Println("Function invoke in function")

	powerOfTwo := func(i int) int {
		return i * i
	}

	executableFunction(powerOfTwo, 2)

	returnFunction("shamim")()

	res := returnFunction1("Md")("Shamim")
	fmt.Println(res)

	res1 := returnFunction1("Md")
	fmt.Println(res1("Shamim"))

	recursionLoop(5)

	fmt.Println("Fact ====> ", fact(7))

}

func executableFunction(myFunc func(int) int, x int) {
	fmt.Println(myFunc(x))
}

func returnFunction(str string) func() {
	return func() {
		fmt.Println("Hello " + str + " from return function.")
	}
}

func returnFunction1(str string) func(str1 string) string {
	return func(str1 string) string {
		return str + " -- " + str1
	}
}

func arithmeticOperation(num int, num2 int) (sum int, divide float64, multiple int, sub int) {

	sum = num + num2

	yes, _ := checkLessThanOne(num2)

	divide = 0

	if !yes {
		divide = float64(num) / float64(num2)
	}

	multiple = num * num2
	sub = num - num2

	return
}

func arithmeticOperation1(num int, num2 int) (int, float64, int, int) {

	sum := num + num2

	yes, _ := checkLessThanOne(num2)

	divide := 0.0

	if !yes {
		divide = float64(num) / float64(num2)
	}

	multiple := num * num2
	sub := num - num2

	return sum, divide, multiple, sub
}

func recursionLoop(len int) {

	if len <= 0 {
		return
	}

	fmt.Println(len)
	recursionLoop(len - 1)
	fmt.Println(len)
}

func fact(n int) int {

	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}
