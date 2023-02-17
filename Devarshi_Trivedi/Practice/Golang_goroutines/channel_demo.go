package main

import (
	"fmt"
	"sync"
)

func channel_bid() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	kh := make(chan int, 5)
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup, kh chan int) {
		fmt.Println(<-ch)
		fmt.Println(<-kh)
		fmt.Println(<-kh)

		wg.Done()
	}(ch, wg, kh)
	go func(ch chan int, wg *sync.WaitGroup, kh chan int) {
		ch <- 42
		kh <- 10
		kh <- 22
		wg.Done()
	}(ch, wg, kh)
	wg.Wait()
}
