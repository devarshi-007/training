package module4

import (
	"fmt"
	"time"
)

const (
	defFormat     = "15:04:05"
	defDateFormat = "2 Jan 2006"
	defInSec      = 5
	nano          = 1000000000.0
)

func timeStamp() {
	first := time.Now()

	fmt.Printf("It is currently %v\n", first.Format(defFormat))

	time.Sleep(defInSec * nano)

	second := time.Now()

	fmt.Printf("It is now %v\n", second.Format(defFormat))

	elapsed := time.Since(first)

	fmt.Printf("%.0f %.0f %.0f\n", elapsed.Seconds(), elapsed.Hours(), elapsed.Minutes())

	future := first.AddDate(0, 5, -5)

	fmt.Println(future.Format(defDateFormat))

	futureHour := second.Add(5 * time.Hour)

	fmt.Println(futureHour.Format(defFormat))

	fmt.Println(time.Until(futureHour).Hours())

}
