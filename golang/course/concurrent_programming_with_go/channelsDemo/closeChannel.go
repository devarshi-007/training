package channelsdemo

import (
	"bytes"
	"fmt"
	"runtime/debug"
	"sync"
)

func closeChannelsDemo() {

	wg := &sync.WaitGroup{}

	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {

		gr := bytes.Fields(debug.Stack())[1]
		fmt.Println("id: ", string(gr))

		fmt.Println("message received 1: ", <-ch)
		// close(ch)
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		}
		// fmt.Println(<-ch)
		// ch <- 5
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 42
		fmt.Println("message sent")
		close(ch)
		//ch <- 50
		//fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}