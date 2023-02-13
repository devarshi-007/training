package main //main package

import "fmt" //fmt package is imported here

/* you can also import multi package using this

method 1:-
import "abc"
import "xyz"

method 2:-
import(
	"abc"
	"xyz"
*/

func hello_world() { //main function

	fmt.Println("Hello, World!") //print the hello world using fmt
	/* you can also use build in println
	println("Hello, WOrld!")
	*/
}

/* Note :-

Golang use the all imported or the declared things
so, if you imported a package but  dosen't use it, it gives you an error */
