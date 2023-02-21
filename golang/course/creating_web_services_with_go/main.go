package main

import (
	"fmt"
	handlinghttprequestwebservices "main/handlingHttpRequest.webservices"
)

func main() {
	fmt.Println("Welcome to Creating Web Services With Go 1")

	// handlinghttprequests.ManageHandler()

	handlinghttprequestwebservices.ManageHandler()
}
