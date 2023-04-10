package channelsdemo

import (
	"fmt"
	"math/rand"
	"sync"
)

func forLoop() {
	wg := &sync.WaitGroup{}

	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {

		i := 1
		for msg := range ch{			
			fmt.Println("message received 1: ", msg, i)
			i++
		}
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- rand.Intn(100)
		}
		fmt.Println("message sent")
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
