package main

import "fmt"

func panic_recover() {
	defer Rrecover()
	defer fmt.Println("HI")
	// panic("I am Panic...........")
	fmt.Println("oorah")

}

func Rrecover() any {
	a := recover()
	fmt.Println("Bye Panic....")
	return a
}
