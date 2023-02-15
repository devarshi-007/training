package main

import "fmt"

func array() {
	var arr [3]int
	arr[0] = 1
	// arr[1] = 1
	// arr[2] = 1
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 2}
	fmt.Println(arr2)

	slice := []int{1, 20}
	s2 := arr[:2]
	fmt.Println(slice)
	fmt.Println(s2)

}
