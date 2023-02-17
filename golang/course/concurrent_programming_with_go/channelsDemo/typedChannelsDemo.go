package channelsdemo

import (
	"fmt"
	"sync"
)

func typedChannelsDemo() {

	wg := &sync.WaitGroup{}

	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println("message received 1: ", <-ch)
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 42
		fmt.Println("message sent")
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
