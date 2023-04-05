package goroutines

import (
	"fmt"
	"main/shared"
	"time"
)

func Main() {

	for i := 0; i < 20; i++ {
		id := shared.Rnd.Intn(shared.LenOfBooks()) + 1

		fmt.Printf("\nquery no: %d\n", i+1)

		go func(id int) {
			if b, ok := shared.QueryCache(id); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
		}(id)

		go func(id int) {
			if b, ok := shared.QueryDatabase(id); ok {
				fmt.Println("from  database")
				fmt.Println(b)
			}
		}(id)

		fmt.Printf("Book not found with id : %v\n", id)
		time.Sleep(150 * time.Millisecond)
	}

	time.Sleep(1 * time.Second)

	fmt.Println("\nend of goroutines")
}
