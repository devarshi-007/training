package syncpackage

import (
	"fmt"
	"main/shared"
	"sync"
)

func syncWithRWMutex() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	for i := 0; i < maxTest; i++ {
		id := shared.Rnd.Intn(shared.LenOfBooks()) + 1

		fmt.Printf("\nquery no: %d\n", i+1)
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := QueryCacheWithRwMutex(id, m); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := QueryDatabaseWithRwMutex(id, m); ok {
				fmt.Println("from  database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)

		fmt.Printf("Book not found with id : %v\n", id)
	}

	wg.Wait()
}
