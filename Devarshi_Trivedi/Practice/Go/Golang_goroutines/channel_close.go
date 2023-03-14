package main

import (
	"fmt"
	"sync"
)

func channel_close() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	//kh := make(chan int, 5)
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		close(ch)         //if channel is recive only you cannot close the channel
		fmt.Println(<-ch) //it gives 0 if it empties
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		//close(ch)
		//ch <- 42 cannot do
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
