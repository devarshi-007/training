package main

import "fmt"

func channel_select() {
	c1 := make(chan int)
	c2 := make(chan string)

	select {
	case i := <-c1:
		fmt.Println(i)
	case c2 <- "hello":
		c2 <- "hello"
		fmt.Println(c2)
	default:
		fmt.Println("none")
	}
}
