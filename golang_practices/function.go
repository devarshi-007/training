package main

import (
	"fmt"
)

type User struct {
	ID  int
	fnm string
}

var (
	users  []*User
	nextID = 1
)

func functionex() {
	port := 3000
	err := startserver(port)
	fmt.Println(err)
}
func startserver(port int) error {
	fmt.Println("starting ...")
	fmt.Println("started on port", port)
	return nil
}

func getUser() []*User {
	return users
}
