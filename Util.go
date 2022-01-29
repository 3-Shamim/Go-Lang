package main

import "errors"

func checkLessThanOne(number int) (bool, error) {

	if number > 0 {
		return false, nil
	}

	return true, errors.New("given number is less than 1")
}
