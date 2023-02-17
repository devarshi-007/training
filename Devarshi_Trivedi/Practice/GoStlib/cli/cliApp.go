package cli

import (
	"fmt"
	"runtime"
)

func main() {
	Version()
}

func Version() {
	fmt.Println("Current Version ", runtime.Version(), "running in", runtime.GOOS)
}
