package main

import "fmt"

func variables() {
	//there are 2 ways to declare a variable
	var a int
	a = 5

	//method 2
	b := 10

	//here it take b type implicitly

	/*one main disadvantage is :-
	for arithmatic operation of different types method 2 var can implicitly change
	but as for method1 it have to  explicitly
	*/

	println(a, b)

	//or you can do this

	var (
		d = 5
		e = 6
	)

	println(d, e)
}

func complexian() {
	var i complex64 = complex(3, 2)
	fmt.Println(i)
	k, l := real(i), imag(i)
	fmt.Println(k, l)
}
