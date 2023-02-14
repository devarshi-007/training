package main

import "fmt"

type details struct {
	x, y int
}

func formatting_data_types() {
	p := details{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%b\n", p)   //binary
	fmt.Printf("%c\n", 33)  //aasci
	fmt.Printf("%x\n", 455) //hex
}
