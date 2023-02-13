package module2

import "fmt"

func mapDataType() {
	m := map[string]int{"key": 0}
	fmt.Println(m, m["key"])

	m["key"] = 1
	fmt.Println(m, m["key"])

	delete(m, "key")
	fmt.Println(m, m["key"])

}
