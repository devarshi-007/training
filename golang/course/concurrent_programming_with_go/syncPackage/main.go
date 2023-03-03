package syncpackage

import (
	"fmt"
)

const (
	maxTest = 10
)

func Main() {

	syncWithMutex()

	syncWithRWMutex()

	fmt.Println("\nend of syncPackage")
}
