package module1

import (
	"fmt"
	"os"
	"strconv"
)

func runtimeArgs(isShow bool) {
	args := os.Args
	convertableArgs := []int{}

	fmt.Printf("file name: %s", args[0])

	sum := 0

	for _, v := range args {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("%s can not convert to int\n", v)
			continue
		}
		sum += i
		convertableArgs = append(convertableArgs, i)
	}

	if isShow {
		fmt.Printf("sum of %v is %d\n", convertableArgs, sum)
	} else {
		fmt.Printf("sum of is %d\n", sum)
	}

}
