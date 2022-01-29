package main

import "fmt"

func main() {

	x := [5]int{1, 2, 3, 4, 5}

	fmt.Println(x)

	s := x[1:3]

	fmt.Println(s)

	fmt.Println(x[1:])
	fmt.Println(x[:3])

	b := append(s, 10)
	s = append(s, 10)
	fmt.Println(b)

	a := make([]int, 2) // set capacity 2 this array
	fmt.Println("a ===========", a)

	fmt.Println(len(x), cap(x), len(s), cap(s))

}
