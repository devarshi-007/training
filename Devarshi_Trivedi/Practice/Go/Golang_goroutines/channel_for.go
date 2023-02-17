package main

import (
	"fmt"
	"sync"
)

func channel_for() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	//kh := make(chan int, 5)
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
