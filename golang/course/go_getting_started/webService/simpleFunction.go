package webService

import (
	"errors"
	"fmt"
)

func webtest() {
	fmt.Println("Hello world")
	port := 3000
	_, err := startWebServer(port, 5)
	fmt.Printf("\n%v \n%T", err, err)
}

func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")
	fmt.Println("Server started on port", port)
	return port, errors.New("something went wrong")
}

func giveDummy() {
	u1 := User{
		ID:        3,
		FirstName: "abc",
		LastName:  "xyz",
		Done:      OurCustomClass{IsAllGood: true},
	}

	u2 := User{
		ID:        4,
		FirstName: "abc",
		LastName:  "xyz",
		Done:      OurCustomClass{IsAllGood: false},
	}

	fmt.Println(u1, u2)
}
