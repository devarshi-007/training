package main

import "fmt"

func arary() {
	//ways of decleration

	//1
	var arr [10]int
	arr[0] = 5
	fmt.Println(arr)

	//2
	arr1 := [10]int{1, 2, 4}
	fmt.Println(arr1)
	fmt.Println(arr1[:5])
}
