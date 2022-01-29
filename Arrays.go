package main

import "fmt"

func main() {

	fmt.Println("Arrays")

	var arr [5]int
	arr[0] = 1
	fmt.Println(arr)

	arr1 := []int{1, 2, 3}
	fmt.Println(arr1)

	arr2 := []string{"One", "Two", "Three"}
	fmt.Println(arr2)

	fmt.Println("Length of arr 2:", len(arr2))

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	for index, value := range arr2 {
		fmt.Println(index, value)
	}

	arr2D := [][]int{{1, 2, 3}, {1, 2, 3}}

	fmt.Println(arr2D[0][2])

	for i := 0; i < len(arr2D); i++ {

		for j := 0; j < len(arr2D[i]); j++ {
			fmt.Printf("%d ", arr2D[i][j])
		}
		fmt.Println()

	}

}
