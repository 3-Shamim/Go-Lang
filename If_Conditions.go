package main

import "fmt"

func main() {

	age := 5

	if age >= 18 {
		fmt.Println("You can do this.")
	}

	if 5 < 6 {
		fmt.Println("Hello if")
	} else if 4 > 2 {
		fmt.Println("Hello else if")
	} else if 4 > 2 || 3 == 3 {
		fmt.Println("Hello else if 2")
	} else {
		fmt.Println("Hello else")
	}

}
