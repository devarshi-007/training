package module5

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func stringSplit() {

	str := "this is something different"

	// strCollection := strings.Split(str, " ")
	// strCollection := strings.SplitAfter(str, " ")
	strCollection := strings.SplitAfterN(str, " ", 2)

	fmt.Printf("_%s_ _%s_ %v %d %T", strCollection[0], strCollection[1], strCollection, len(strCollection), strCollection)

	file, _ := os.Open("module2/fixedCsv.csv")

	scanner := bufio.NewScanner(file)

	lineNo := 1

	fmt.Println()

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.SplitAfter(line, ",")

		fmt.Printf("-- record [%d] --\n", lineNo)

		for i := range items {
			fmt.Println(strings.TrimSpace(items[i]))
		}

		lineNo++
	}
}
