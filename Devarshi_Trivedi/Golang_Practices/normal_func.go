package main

import "fmt"

func call_funcs() {
	normal_func()
	func_with_arg(5)
	r1 := func_are_and_ret(5, 6)
	fmt.Println(r1)
	r2 := without_arg_ret()
	fmt.Println(r2)
}
func normal_func() {
	fmt.Println("hi")
}

func func_with_arg(a int) {
	fmt.Println(a)
}

func func_are_and_ret(a, b int) int {
	return a + b
}

func without_arg_ret() int {
	return 5
}
