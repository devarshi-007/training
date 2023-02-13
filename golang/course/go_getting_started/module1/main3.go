package module1

import "fmt"

const (
	mine  = "a"
	yours = iota << 2
	double
)

const (
	doubledouble = iota
	not
)

func constdemo() {
	// const c = 10
	// fmt.Print(c + 1.2)

	// const c int= 10
	// fmt.Print(c + 1.2)

	const pi float32 = float32(22) / 7

	fmt.Println(pi, first, last, mine, yours, double, doubledouble, not)

}
