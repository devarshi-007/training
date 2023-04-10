package syncpackage

import (
	"fmt"
	"main/shared"
	"sync"
	"time"
)

func QueryCacheWithRwMutex(id int, m *sync.RWMutex) (shared.Book, bool) {
	m.RLock()
	b, ok := shared.Cache[id]
	fmt.Println("cash res", b, ok, id, shared.Cache[id])
	m.RUnlock()
	return b, ok
}

func QueryDatabaseWithRwMutex(id int, m *sync.RWMutex) (shared.Book, bool) {
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
