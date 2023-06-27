package concurrencyless

import (
	"fmt"
	"main/shared"
	"time"
)

func Main() {
	for i := 0; i < 10; i++ {
		id := shared.Rnd.Intn(shared.LenOfBooks()) + 1

		if b, ok := shared.QueryCache(id); ok {
			fmt.Println("from cache")
			fmt.Println(b)
			continue
		}

		if b, ok := shared.QueryDatabase(id); ok {
			fmt.Println("from  database")
			fmt.Println(b)
			continue
		}

		fmt.Printf("Book not found with id : %v", id)
		time.Sleep(150 * time.Millisecond)
	}

	fmt.Println("\nend of general secenario")
}
