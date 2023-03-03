package module2

import "fmt"

const (
	InfoColor        = "\033[1;34m%s\033[0m"
	WarningColor     = "\033[1;33m%s\033[0m"
	RecommandadColor = "\033[1;32m%s\033[0m"
	ErrorColor       = "\033[1;31m%s\033[0m"
)

func formatOutput() {
	fmt.Printf("|%-5.2f|%-5.2f|\n", 10.22, -100.0)
	fmt.Printf("|%5.2f|%5.2f|\n", 10.22, -100.0)
	fmt.Printf("|%.2f|%.2f|\n", 10.22, -100.0)
	fmt.Printf("|%f|%f|\n", 10.22, -10.0)
	fmt.Printf("|%7s|%7q|\n", "hello", "hii")

	fmt.Printf(RecommandadColor, fmt.Sprintf("\n Message: %s\n", "this is a message"))

	fmt.Printf("%v %b _%c_ %x %o %d", []int8{1, 2}, 10, 41, 42, 12, 12)
}
