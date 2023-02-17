package main

type shape interface {
	Area() int
}

type circle struct {
	radius int
}

type rect struct {
	height int
	width  int
}

func (c circle) Area() int {
	return 3 * c.radius * c.radius
}

func (r rect) Area() int {
	return r.height * r.width
}

func interfaces() {
	o1 := circle{2}
	o2 := rect{2, 3}
	shapes := []shape{o1, o2}

	for _, x := range shapes {
		println(x.Area())
	}
}
