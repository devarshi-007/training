package module1

import (
	"fmt"
)

const myint int = 5

func pointerDemo() {

	var name *string = new(string)
	*name = "hello"

	// var otherstr string = "hello"
	// var name *string = &otherstr

	fmt.Println(*name, name, &name, *&name)
	fmt.Printf("*name == *&name: %t %v", name == *&name, new(string))

	// fmt.Println( /* *otherstr, */ otherstr, &otherstr)
	// fmt.Printf("\n&name: (%T, %v) otherstr: (%T, %v) name: (%T, %v)\n", &name, &name, otherstr, otherstr, name, name)
}
