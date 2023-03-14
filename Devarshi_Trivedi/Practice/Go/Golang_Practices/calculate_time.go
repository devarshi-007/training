package main

import (
	"fmt"
	"time"
)

func calculate_time() {
	first := time.Now()
	fmt.Printf("Time is: %v \n", first.Format("15:04:05"))
	time.Sleep(3 * time.Second)
	second := time.Now()
	fmt.Printf("Now time is: %v", second.Format("15:04:05"))
}

func elapsed_time() {
	today := time.Now()
	fmt.Printf("It is currently %v\n", today.Format("Monday, Jan 2 2006"))
	startDate := time.Date(2023, 02, 9, 03, 27, 00, 00, time.UTC)
	fmt.Printf("that day: %v\n", startDate.Format("Monday, 2 Jan 2006"))
	elapsed := time.Since(startDate)
	fmt.Printf("time is %.2v %v %v\n", elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds())
}

func other_way() {
	right := time.Now()
	next := right.AddDate(0, 6, 0)
	nextf := right.Add(6 * time.Hour)
	fmt.Printf("Date of next %v %v", next.Format("Monday, 5 Jan 2006"), nextf.Format("Monday, 2 Jan 2006"))
}

func bedtime() {
	bedtime := time.Date(2023, 2, 10, 23, 00, 00, 00, time.Local)
	fmt.Printf("There is %.0f hours until bedtime\n", time.Until(bedtime).Hours())
}
