package main

import "fmt"

func loop() {
	var i int
	for i < 10 {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
		if i == 2 {
			continue
		}
		println("..")
	}
}
