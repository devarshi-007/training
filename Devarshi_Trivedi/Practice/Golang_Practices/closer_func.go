package main

import "fmt"

func normal_closer_func() {
	//fmt.Println("this is closer func")
	d := my_f_wr()
	fmt.Println(d())
}

func my_f_wr() func() int {
	return func() int {
		return 5
	}
}
