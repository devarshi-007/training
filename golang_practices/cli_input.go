package main

import "fmt"

func cli_input() {
	var name string
	fmt.Println("What's your name?")
	input, _ := fmt.Scanf("%s", &name)
	fmt.Printf("hello %s input is %d\n", name, input)

}
