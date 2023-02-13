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
	fmt.Printf("Hello %v", text)

	fmt.Println("\nend of the file")
}
