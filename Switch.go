package main

import "fmt"

func main() {

	ans := 5.5

	switch ans {
	case 5, 5.5:
		fmt.Println("It's 5")
	case 10:
		fmt.Println("Wow It's 10")
	default:
		fmt.Println("Default print")
	}

	switch {
	case ans >= 5 && ans < 6:
		fmt.Println("It's 5")
	default:
		fmt.Println("Default print")
	}

}
