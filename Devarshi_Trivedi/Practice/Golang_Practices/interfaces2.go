package main

type Taker struct {
	name string
	id   int
}

type Bird struct {
	name string
	id   int
}

func (u Taker) getName() string {
	return u.name
}

func (b Bird) getName() string {
	return b.name
}

type gn interface {
	getName() string
}

func start() {
	u := Taker{"Ben10", 1}
	b := Bird{"Sparrow", 2}
	group := []gn{u, b}

	for _, x := range group {
		println(x.getName())
	}
}
