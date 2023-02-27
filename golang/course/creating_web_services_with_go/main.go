package main

import (
	"fmt"
	filehandling "main/fileHandling"
)

func main() {
	fmt.Println("Welcome to Creating Web Services With Go 1")

	// handlinghttprequests.ManageHandler()

	// handlinghttprequestwebservices.ManageHandler()

	// httpRequestWithDatabase.ManageHandler()

	filehandling.HandleRoute("/mid")
}
