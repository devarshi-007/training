package main

import (
	"log"
	"time"
)

func calc_time() {
	startTime := time.Now()
	time.Sleep(2 * time.Second)
	writeTime(startTime)
}

func writeTime(startTime time.Time) {
	elapsed := time.Since(startTime)
	log.Printf("took %s", elapsed)
}
