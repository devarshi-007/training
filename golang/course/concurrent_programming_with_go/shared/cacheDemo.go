package shared

import (
	"math/rand"
	"time"
)

var Cache = map[int]Book{}

func LenOfBooks() int {
	return len(Booklist)
}

func QueryCache(id int) (Book, bool) {
	b, ok := Cache[id]
	return b, ok
}

func QueryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range Booklist {
		if b.ID == id {
			Cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}

var Rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
