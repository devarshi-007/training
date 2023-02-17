package main

import (
	"fmt"
	"time"
)

func timer() {
	t := time.Now()
	fmt.Print(t, "\n")
	fmt.Print(t.Day(), "-", t.Year(), "\n")
	fmt.Println(t.Format("2 January 2006"))

	startDate := time.Date(2033, 12, 12, 00, 00, 00, 00, time.UTC) //
	fmt.Println(startDate.Format("2 January 2006"))
}
