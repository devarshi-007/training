package module4

import (
	"fmt"
	"time"
)

func dateAndTime() {
	t := time.Now()

	fmt.Print(t.Year(), t.Month(), t.Day(), t.Weekday())

	fmt.Println(t.Format(time.ANSIC), t.Format(time.RFC3339), t.Format(time.UnixDate))

	fmt.Println(t.Format(time.Kitchen))

	fmt.Println(t.Format("January, 2 2006 (Mon) | 15:04:05"))

	// startDate := time.Date(year, month, date, hower, minutes, seconds, nanoseconds, time format)
	startDate := time.Date(2000, 4, 27, 15, 0, 0, 0, time.Now().Local().Location())

	fmt.Println(startDate)
}
