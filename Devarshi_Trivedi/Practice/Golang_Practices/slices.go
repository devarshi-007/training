package main

import "fmt"

func slices() {

	s := []int{1, 2, 3, 4}

	//or

	arr := [3]int{1, 2, 3}
	s1 := arr[:3]

	s = append(s, 2, 3, 4, 5, 6, 7)
	fmt.Println(s, s1)

	//or

	s3 := make([]int, 3, 10) //where []int =slice type, 3= len, 10 = cap
	fmt.Println(s3)

}
