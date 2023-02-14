package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	//wg := sync.WaitGroup{}
	//wg := &sync.WaitGroup{} //we gonna passing this around in app , so we can pass pointer
	count := 1
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		go func(id int) {
			if b, ok := queryCache(id); ok {
				fmt.Println("Count: ", count)
				count++
				fmt.Println("from cache")
				fmt.Println(b)
			}
		}(id)
		go func(id int) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("Count: ", count)
				count++
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(id)
		//fmt.Println("Book not found of id", id)

		//Right way of doing this is waitgroups
		time.Sleep(150 * time.Millisecond)
		//DEMO WORKING ONLY BECAUSE OF THIS, REMOVE THIS AND IT DOSEN'T WORK
	}
	//time.Sleep(2 * time.Second)
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}
