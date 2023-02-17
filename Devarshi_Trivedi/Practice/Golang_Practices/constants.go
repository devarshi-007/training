package main

const (
	a = iota
	b
)

const (
	c = iota + 1
	d
)

func consta() {
	const e = 5
	println(a, b, c, d, e)
}
