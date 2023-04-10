package main

import (
	"fmt"
	channelsdemo "main/channelsDemo"
)

func main() {
	fmt.Println("start of concurrent programming with go")

	// fmt.Println( reflect.TypeOf(time.Now()).Kind())

	// concurrencyless.Main()

	// goroutines.Main()

	// syncpackage.Main()

	channelsdemo.Main()
}
