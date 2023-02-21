package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache1 = map[int]Book{}
var rnd1 = rand.New(rand.NewSource(time.Now().UnixNano()))

func waitgroup() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup) {
			if b, ok := queryCache1(id); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)
		go func(id int, wg *sync.WaitGroup) {
			if b, ok := queryDatabase1(id); ok {
				fmt.Println("from database")
				// cache[id] = b
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)
		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()
}

func queryCache1(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase1(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			return b, true
		}
	}

	return Book{}, false
}
