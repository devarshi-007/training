package main

import (
	"fmt"
	socketdemo "main/socketDemo"
)

func main() {
	fmt.Println("Welcome to Creating Web Services With Go 1")

	// handlinghttprequests.ManageHandler()

	// handlinghttprequestwebservices.ManageHandler()

	// httpRequestWithDatabase.ManageHandler()

	// filehandling.HandleRoute("/mid")

	socketdemo.StartListening()

}
