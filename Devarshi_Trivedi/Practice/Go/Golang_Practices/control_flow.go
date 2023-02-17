package main

func control_flow() {
	i := 5

	if i == 0 {
		println("0")
	} else if i < 5 {
		println("1-4")
	} else {
		println("5 or more")
	}

	switch i {
	case 1:
		println("1")
	case 2:
		println("2")
	case 5:
		println("5")
		fallthrough
	case 6:
		println("6")
	default:
		println("default")
	}
}
