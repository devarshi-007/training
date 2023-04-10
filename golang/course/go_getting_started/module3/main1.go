package module3

import "fmt"

func controllStatement() bool {

	var start int = 0
	var end int = 5
	var temp int = start
	var arr [5]int

	// while
	for temp <= end {
		temp++

		if temp == 4 {
			break
		}
		if temp == 2 {
			continue
		}

		fmt.Println(temp)

	}

	// simple for
	temp = start
	for ; temp < end; temp++ {
		fmt.Println(temp)
	}

	// infinite loop
	temp = start
	for {
		if temp == end {
			break
		}
		arr[temp] = temp * 2
		temp++
	}

	// loop over collections
	for i, v := range arr {
		fmt.Println(i, v)
	}

	for i := range arr {
		fmt.Println(i)
	}

	// if demo
	var first = "hello"
	var second = first
	var third = "hello"
	if &first == &second {
		fmt.Println("&first == &second")
	} else if first == third {
		fmt.Println("first == third")
	} else {
		fmt.Println("first == second != third")
	}

	//switch
	switch 123 {
	case 124:
		fmt.Println("124")
	case 125:
		fmt.Println("125")
		fallthrough
	case 123:
		fmt.Println("123")
		fallthrough
	default:
		println("default")
	}

	switch 0 {
	case 1:
		return true
	default:
		return false
	}

}
