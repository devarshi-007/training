package main

import "fmt"

func point() {
	//2 ways

	//way of sages
	var i *int = new(int)
	*i = 42
	println(*i)

	//2nd

	var j *int
	data := 42
	j = &data //or *j = data
	println(*j)

	//3rd

	kate := 42
	pointy := &kate
	fmt.Println(*pointy)
}
