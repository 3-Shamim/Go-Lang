package main

import "fmt"

func main() {

	fmt.Printf("Hello %T %v\n", 10, 10)
	var str string = fmt.Sprintf("Hello %s", "world")
	fmt.Println(str)

}
