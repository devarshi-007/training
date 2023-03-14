package main

import "fmt"

func formatting_output() {
	fmt.Print("Hi")
	fmt.Print("Hello\n")
	age := 10
	fmt.Print("Age is", age, "Hi\n")
	var out, _ = fmt.Print("Age is", age)
	fmt.Print("Bytes Written: ", out)

	fmt.Println("\n")

	var age1 int = 17
	var name string = "Tom"

	fmt.Printf("Name is %v and age is %d\n", name, age1)

	var pi float32 = 3.14
	fmt.Printf("%f\n", pi)
	fmt.Printf("%.2f\n", pi)

	a, b, c := 1.100, 12.12, 13.12323
	fmt.Printf("%f %f %f\n", a, b, c)
	fmt.Printf("|%-2.2f|%-1.2f|%-1.2f|\n", a, b, c)

	fmt.Printf("|%7qs|%7s|%7s|\n", "foo", "baar", "took")
	fmt.Printf("|%-7s|%-7s|%-7s|\n", "foo", "baar", "took")

	output := fmt.Sprintf("Hi")
	fmt.Println(output)

}
