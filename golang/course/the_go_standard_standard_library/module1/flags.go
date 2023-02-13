package module1

import (
	"flag"
	"fmt"
)

func flagDemo() {

	flgPtr := flag.Bool("sum", true, "to sum or show")

	flag.Parse()

	fmt.Println(*flgPtr)

	runtimeArgs(*flgPtr)
}
