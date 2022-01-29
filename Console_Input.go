package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	//input1 := scanner.Text() 2nd time read
	numb, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		fmt.Println("============", numb)
	}

	fmt.Println(input, numb)

}
