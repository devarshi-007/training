package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf("Title:\t\t%q\n"+
		"Author:\t\t%q\n"+
		"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	Book{
		ID:     1,
		Title:  "aaa",
		Author: "xxx",
	},
	Book{
		ID:     2,
		Title:  "bbb",
		Author: "yyy",
	},
	Book{
		ID:     3,
		Title:  "ccc",
		Author: "zzz",
	},
	Book{
		ID:     4,
		Title:  "ddd",
		Author: "www",
	},
	Book{
		ID:     5,
		Title:  "fff",
		Author: "vvv",
	},
}
