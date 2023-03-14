package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	{
		ID:            1,
		Title:         "Bhedi Tapu",
		Author:        "I.K.Vijalyvala",
		YearPublished: 2006,
	},
	{
		ID:            2,
		Title:         "Bhedi Tapu 2",
		Author:        "I.K.Vijalivala",
		YearPublished: 2007,
	},
	{
		ID:            3,
		Title:         "Bhedi Tapu3",
		Author:        "I.K.Vijalyvala",
		YearPublished: 2010,
	},
	{
		ID:            4,
		Title:         "Bhedi Tapu 4",
		Author:        "I.K.Vijalivala",
		YearPublished: 2011,
	},
	{
		ID:            5,
		Title:         "Bhedi Tapu5",
		Author:        "I.K.Vijalyvala",
		YearPublished: 2017,
	},
	{
		ID:            6,
		Title:         "Bhedi Tapu 6",
		Author:        "I.K.Vijalivala",
		YearPublished: 2017,
	},
	{
		ID:            7,
		Title:         "Bhedi Tapu7",
		Author:        "I.K.Vijalyvala",
		YearPublished: 2018,
	},
	{
		ID:            8,
		Title:         "Bhedi Tapu 8",
		Author:        "I.K.Vijalivala",
		YearPublished: 2019,
	},
	{
		ID:            9,
		Title:         "Bhedi Tapu9",
		Author:        "I.K.Vijalyvala",
		YearPublished: 2023,
	},
	{
		ID:            10,
		Title:         "Bhedi Tapu 10",
		Author:        "I.K.Vijalivala",
		YearPublished: 2023,
	},
}
