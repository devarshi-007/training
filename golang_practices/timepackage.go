package main

import (
	"fmt"
	"time"
)

func timepackage() {

	t := time.Now()
	fmt.Print(time.Now(), "\n")

	fmt.Print(t.Year(), "\n")
	fmt.Print(t.Month(), "\n")

	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Format(time.UnixDate))
	fmt.Println(t.Format(time.Kitchen))

	startdate := time.Date(2000, 12, 20, 9, 0, 9, 0, time.UTC)
	elapsed := time.Since((startdate))

	//fmt.Printf("Years:%v\n", elapsed.Year())

	fmt.Printf("Hours :%v,Minute:%v,second:%v\n", elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds())

}
