package module3

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

// go tool trace trace.out

func tranceLoggerManager() {
	f, err := os.Create("module3/trace.out")

	if err != nil {
		log.Fatalf("cann't create file :: %v\n", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("cann't close file :: %v\n", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("cann't start trace :: %v\n", err)
	}

	defer trace.Stop()

	traceLogger()
}

func traceLogger() {
	first := rand.Intn(100)
	second := rand.Intn(100)

	time.Sleep(2 * time.Second)

	fmt.Printf("Result of two is %d\n", first*second)
}
