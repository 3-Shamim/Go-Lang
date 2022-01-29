package main

import (
	"fmt"
)

func main() {

	//var name string = "Hello World!"
	var name string
	name = "I will be removed."
	name = "Hello World!"

	fmt.Println(name)

	//var number int16 = -260
	//var number int = 260
	var number uint16 = 260
	number += 5
	fmt.Println(number)

}
