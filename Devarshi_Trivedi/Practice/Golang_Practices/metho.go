package main

import "fmt"

type User struct {
	name string
	id   int
}

func (u User) AddUser(s string) User {
	u.name = s
	u.id = u.id + 1
	return u
}

func metho() {
	u := User{}
	fmt.Println(u.AddUser("me"))
}
