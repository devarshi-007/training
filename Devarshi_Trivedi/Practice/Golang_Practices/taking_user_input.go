package main

import "fmt"

func taking_user_input() {
	var i int
	fmt.Scanln(&i) //The Scanln is similar to Scan, but stops scanning at a newline.
	fmt.Println(i)

	var j int
	fmt.Scan(&j)
	fmt.Print(j, "\n")

	var k int
	fmt.Scanf("%d", &k)
	fmt.Printf("%d\n", k)

}
