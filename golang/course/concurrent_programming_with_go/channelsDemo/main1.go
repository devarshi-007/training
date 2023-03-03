package channelsdemo

import (
	"fmt"
	"main/shared"
	"sync"
	"time"
)

func queryDatabase(id int) (shared.Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range shared.Booklist {
		if b.ID == id {
			return b, true
		}
	}

	return shared.Book{}, false
}

func syncWithRWMutexAndChannel() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	CacheCh := make(chan shared.Book)
	dbCh := make(chan shared.Book)

	for i := 0; i < maxTest; i++ {
		id := shared.Rnd.Intn(8) + 1

		wg.Add(2)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- shared.Book) {
			if b, ok := shared.QueryCache(id); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, CacheCh)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- shared.Book) {
			if b, ok := queryDatabase(id); ok {
				m.Lock()
				shared.Cache[id] = b
				m.Unlock()
				ch <- b
			} else{
				ch <- shared.Book{ID: id}
			}
			wg.Done()
		}(id, wg, m, dbCh)

		go func(cacheCh, db <-chan shared.Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(CacheCh, dbCh)

		time.Sleep(120 * time.Millisecond)
		//fmt.Printf("Book not found with id : %v\n", id)
	}

	wg.Wait()
}
