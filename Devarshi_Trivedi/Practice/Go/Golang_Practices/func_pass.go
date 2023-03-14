package main

import "fmt"

func hello(xz func(int) int) func() int {
	x := xz(10)
	return func() int {
		fmt.Println("Iteration 1", x)
		x := x + 20
		fmt.Println("Iteration 2", x)
		return x
	}
}

func pass_to_this(i int) int {
	k := i + 20
	return k
}
func func_pass() {
	y := 10
	z := hello(pass_to_this)
	a := z()
	fmt.Println(y)
	fmt.Println(a)
}
