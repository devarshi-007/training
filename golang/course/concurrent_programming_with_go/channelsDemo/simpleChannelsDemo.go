package channelsdemo

import (
	"fmt"
	"sync"
)

func simpleChannelsDemo() {
	wg := &sync.WaitGroup{}

	ch := make(chan int, 2)

	wg.Add(3)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("message received 1: ", <-ch)
		ch <- 0
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		ch <- 50
		fmt.Println("message sent")
		fmt.Println("message received", <-ch, <-ch)
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("message received 2: ", <-ch)
		ch <- 1
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
