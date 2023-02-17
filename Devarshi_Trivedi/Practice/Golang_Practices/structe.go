package main

import "fmt"

type dev struct {
	msg string
}

type me int

func structe() {
	type User struct {
		name string
		id   int
	}

	var u1 User
	u1.name = "N1"
	u1.id = 1

	u2 := User{"N2", 2}

	s1 := dev{"hi"}
	var s2 me
	s2 = 5
	fmt.Println(u1, u2, s1, s2)
}
