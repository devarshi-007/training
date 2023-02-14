package main

import "fmt"

func any_store_var() {
	var k any
	k = 5
	k = "hello"
	k = true
	k = 5.5
	fmt.Println(k)
}
