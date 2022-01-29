package main

import "fmt"

func main() {

	var _map map[string]int = map[string]int{
		"apple": 5,
		"pear":  6,
	}

	var _map1 = map[string]int{
		"apple": 5,
		"pear":  6,
	}

	fmt.Println(_map, _map1)

	var _m = make(map[string]int)
	_m["a"] = 1
	_m["b"] = 2
	_m["c"] = 3

	val, ok := _m["aa"] // ok true if present

	fmt.Println(val, ok)

	fmt.Println(_m)
	fmt.Println("a", _m["a"])

	delete(_m, "a")

	for key, value := range _m {
		fmt.Println(key, value)
	}

}
