package module2

import "fmt"

func arrayAndSlize() {
	var arr [5]int
	arr[3] = 5
	fmt.Println(arr[0])

	copy := arr
	slice := arr[1:4]

	arr = [5]int{1, 2, 1}

	list := []int{1, 2, 3}

	newlist := list

	list = append(list, 4, 5, 6)

	fmt.Println(arr, copy, slice, list, newlist, len(list))

	deleteElm := 2

	list = append(list[:deleteElm], list[deleteElm+1:]...)

	fmt.Println("\n", list)

	list = list[:len(list)-1]

	fmt.Println("\n", list)
}
