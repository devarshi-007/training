package testdemos

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("hello, %v\n", name)
}

func depart(name string) string {
	return fmt.Sprintf("Goodbye, %v\n", name)
}
