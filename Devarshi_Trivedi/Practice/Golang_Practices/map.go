package main

import "fmt"

func mape() {

	m1 := map[string]int{"name1": 1, "name2": 2}
	m2 := make(map[string]int)
	m2["name1"] = 3
	m1["name3"] = 3
	fmt.Println(m1, m2)
	delete(m1, "name1")
	fmt.Println(m1, m2)
	arr := [3]int{1, 2, 3}
	m3 := map[string][]int{}

	m3["arr"] = append(m3["arr"], arr[0])
	m3["arr"] = append(m3["arr"], arr[1])
	m3["arr"] = append(m3["arr"], arr[2])

	fmt.Println(m1, m2, m3)
}
