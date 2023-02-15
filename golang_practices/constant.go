package main

import "fmt"

func constant() {
	const pi = 3.1415
	fmt.Println(pi)

	const c int = 3
	fmt.Println(float32(c) + 1.2)

}
