package main

func various_for() {
	for_with_condition()
	for_with_post_clause()
	for_over_collection()
	infinite_for()
}

func for_with_condition() {
	i := 0
	for i < 5 {
		println(i)
		i++
	}
}

func for_with_post_clause() {
	for i := 0; i < 5; i++ {
		println(i)
	}
	for j := 0; j < 5; {
		println("Hi")
		j++
	}
}
func for_over_collection() {
	arr := [3]int{1, 2, 3}

	for x := range arr {
		println(arr[x])
	}
}
func infinite_for() {
	i := 0
	for {
		if i > 5 {
			break
		}
		println("Hi")
		i++
	}
}
