package main

import (
	"fmt"
	handlinghttprequests "main/handlingHttpRequests"
)

func main() {
	fmt.Println("Welcome to Creating Web Services With Go 1")

	handlinghttprequests.ManageHandler()
}
