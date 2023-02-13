package main

import "fmt"


type Book struct {
	ID int
	Title string
	Author string
	YearPublished int
}

func(b Book) String() string {
	return fmt.Sprintf(
			"Title:\t\t%q\n"+
				"Author:\t\t%q\n"+
					"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
	)
}