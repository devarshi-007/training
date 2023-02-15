package main

import "fmt"

func pointer() {
	var firstname *string = new(string)
	*firstname = "abc"
	fmt.Println(firstname)

	firstnm := "def"
	fmt.Println(firstnm)
	ptr := &firstnm
	fmt.Println(ptr, *ptr)

}
