package main

import "fmt"

func see_defer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
