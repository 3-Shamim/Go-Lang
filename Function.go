package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println(randomNumber(5))

	value, err := ownCheckLessThanOne(1)

	if err == nil {
		fmt.Println("Check Less Than One", value)
	} else {
		fmt.Println("Check Less Than One", err, value)
	}

}

func randomNumber(len int) int {

	var result int

	for i := 0; i < len; i++ {
		result = result*10 + int(rand.Float32()*10+1)
	}

	return result
}

func ownCheckLessThanOne(number int) (bool, error) {

	if number > 0 {
		return true, nil
	}

	return false, errors.New("given number is less than 1")
}
