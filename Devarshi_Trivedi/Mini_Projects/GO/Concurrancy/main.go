package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	//wg := sync.WaitGroup{}
	wg := &sync.WaitGroup{} //we gonna passing this around in app , so we can pass pointer
	m := &sync.RWMutex{}
	// count := 1
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		//wg.Add(1) //method-1
		wg.Add(2) //method-2
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryCache(id, m); ok {
				// fmt.Println("Count: ", count)
				// count++
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		//wg.Add(1) //method-1
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryDatabase(id, m); ok {
				// fmt.Println("Count: ", count)
				// count++
				fmt.Println("from database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		//fmt.Println("Book not found of id", id)

		//Right way of doing this is waitgroups
		time.Sleep(150 * time.Millisecond)
		//DEMO WORKING ONLY BECAUSE OF THIS, REMOVE THIS AND IT DOSEN'T WORK
	}
	//time.Sleep(2 * time.Second)
	wg.Wait() //wait till no activities
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
