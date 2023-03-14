package main

import "time"

func local_or_utc() {
	startTime := time.Date(2023, 2, 10, 18, 15, 00, 00, time.Local)
	endTime := time.Date(2023, 2, 10, 18, 15, 00, 00, time.UTC)
	println(startTime.String())
	println(endTime.String())

	/* Output of Both
	2023-02-10 18:15:00 +0530(diff of utc) IST
	2023-02-10 18:15:00 +0000(diff of utc) UTC
	*/
}
