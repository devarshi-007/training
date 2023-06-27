package shared

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(""+
		"ID: %d\n"+
		"\tTitle: %q\n"+
		"\tAuthor: %q\n"+
		"\tPubliched: %v\n", b.ID, b.Title, b.Author, b.YearPublished)
}

var Booklist = []Book{
	{
		ID:            1,
		Title:         "book<ID>",
		Author:        "book<Author>",
		YearPublished: 2000,
	},
	{
		ID:            2,
		Title:         "book<ID>",
		Author:        "book<Author>",
		YearPublished: 2000,
	},
	{
		ID:            3,
		Title:         "book<ID>",
		Author:        "book<Author>",
		YearPublished: 2000,
	},
	{
		ID:            4,
		Title:         "book<ID>",
		Author:        "book<Author>",
		YearPublished: 2000,
	},
	{
		ID:            5,
		Title:         "book<ID>",
		Author:        "book<Author>",
		YearPublished: 2000,
	},
	{
		ID:            7,
		Title:         "book<ID>",
		Author:        "book<Author>",
		YearPublished: 2000,
	},
	{
		ID:            8,
		Title:         "book<ID>",
		Author:        "book<Author>",
		YearPublished: 2000,
	},
}
