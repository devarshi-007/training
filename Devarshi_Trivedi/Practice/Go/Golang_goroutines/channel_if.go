package main

import (
	"fmt"
	"sync"
)

func channel_if() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	//kh := make(chan int, 5)
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok) //ok->false means channel is closed if true than channel is open
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		close(ch)
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
