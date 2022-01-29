package main

import "fmt"

func main() {

	for i := 0; i < 30; i++ {
		fmt.Println("Hello", i+1)
	}

	x := 3
	for x <= 30 {

		if x == 28 {
			x++
			continue
		}

		if x == 30 {
			break
		}

		fmt.Println("Hi", x)
		x++
	}

	//for true {
	//	fmt.Println("Infinite loop...")
	//}

}
