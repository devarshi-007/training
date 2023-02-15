package main

import "fmt"

func struct1() {
	type user struct {
		ID       int
		firstnm  string
		secondnm string
	}
	var u user
	u.ID = 1
	u.firstnm = "abc"
	fmt.Println(u)

	u2 := user{ID: 1, firstnm: "abc", secondnm: "xyz"}
	fmt.Println(u2)

}
