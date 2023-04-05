package syncpackage

import (
	"main/shared"
	"sync"
	"time"
)

func QueryCacheWithMutex(id int, m *sync.Mutex) (shared.Book, bool) {
	m.Lock()
	b, ok := shared.Cache[id]
	m.Unlock()
	return b, ok
}

func QueryDatabaseWithMutex(id int, m *sync.Mutex) (shared.Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range shared.Booklist {
		if b.ID == id {
			m.Lock()
			shared.Cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return shared.Book{}, false
}
