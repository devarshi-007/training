package module1

import (
	"bufio"
	"fmt"
	"os"
)

func readFile() {
	file, err := os.Open("module1/test.txt")

	if err != nil {
		fmt.Println(err, ":: OS ERROR")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err, ":: SCANNER ERROR")
	}
}
