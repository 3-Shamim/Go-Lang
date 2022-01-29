package main

import "fmt"

func main() {

	x := 5
	y := 6

	val := x > y
	val = x < y

	val = x >= y
	val = x <= y

	val = x != y
	val = x == y

	fmt.Printf("%t\n", val)

}
