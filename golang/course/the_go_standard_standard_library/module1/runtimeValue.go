package module1

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func runtimeValue() {
	fmt.Print(runtime.Version())
	fmt.Print(runtime.Version())

	fmt.Println(runtime.Version())
	fmt.Println(runtime.Version())

	fmt.Printf("\n%v\n", runtime.Version())
	fmt.Printf("%v\n", runtime.Version())

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your name...")
	text, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s", text)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "/quit" {
			fmt.Println("Quitting")
			os.Exit(3)
		} else {
			fmt.Println("You typed" + scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}

	var newName string
	fmt.Println("Enter your new name")
	inps, _ := fmt.Scanf("%s", &newName)
	fmt.Printf("Hello %q %d", newName, inps)

	fmt.Println("\nend of the file")
}
