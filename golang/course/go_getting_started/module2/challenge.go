package module2

import "fmt"

func storeArray() {
	arr := [2]int{}

	m := map[int][]int{}

	m[0] = arr[:]

	fmt.Println(m)
}
